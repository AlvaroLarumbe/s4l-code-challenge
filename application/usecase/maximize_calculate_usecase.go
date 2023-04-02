package usecase

import (
	"github.com/rs/zerolog/log"
	"github.com/yourbasic/graph"
	"s4l-code-challenge/domain/model"
	"sort"
	"time"
)

type MaximizeCalculateUseCase struct {
	statsCalculateUseCase StatsCalculateUseCase
}

func NewMaximizeCalculateUseCase(statsCalculateUseCase StatsCalculateUseCase) MaximizeCalculateUseCase {
	return MaximizeCalculateUseCase{
		statsCalculateUseCase: statsCalculateUseCase,
	}
}

func (u MaximizeCalculateUseCase) Execute(bookingPetitionList []model.BookingPetition) model.Maximized {
	var maximized model.Maximized

	// Sort by check in date
	sort.Sort(model.ByCheckInDate(bookingPetitionList))

	// Assign an integer identifier for each petition
	bookingPetitionGraphMap := make(BookingPetitionMap, len(bookingPetitionList))

	for i, petition := range bookingPetitionList {
		bookingPetitionGraphMap[i+1] = petition
	}

	// Build graph
	g := graph.New(len(bookingPetitionList) + 2) // Adding source and sink vertices
	sinkKey := len(bookingPetitionList) + 1

	for sourceKey, petition := range bookingPetitionGraphMap {
		// First add all edges from SOURCE to every petition, any petition is a valid start point for the graph
		// Costs will be profits in integer cents (147.29 € = 14729)
		g.AddCost(0, sourceKey, int64(bookingPetitionGraphMap[sourceKey].GetProfit()*100))

		// Calculate all edges from each petition to next valid petitions
		availableObjetiveKeys := bookingPetitionGraphMap.FindTargetKeys(petition.GetCheckOut())
		for _, targetKey := range availableObjetiveKeys {
			g.AddCost(
				sourceKey,
				targetKey,
				int64(bookingPetitionGraphMap[targetKey].GetProfit()*100),
			)
		}

		// Add all edges from all petitions to the SINK with cost 0
		g.Add(sourceKey, sinkKey)
	}

	// Calculate topological sort
	topSort, topSortOk := graph.TopSort(g)

	// Begin with weighted directed acyclic graphs algorithm
	// https://en.wikipedia.org/w/index.php?title=Longest_path_problem&oldid=525448430#Weighted_directed_acyclic_graphs
	type path struct {
		value    int64
		vertexes []int
	}

	lengthTo := make([]path, len(bookingPetitionList)+2)

	for _, v := range topSort {
		g.Visit(v, func(w int, cost int64) bool {
			if lengthTo[w].value <= lengthTo[v].value+cost {
				lengthTo[w].value = lengthTo[v].value + cost
				lengthTo[w].vertexes = append(lengthTo[v].vertexes, w)
			}

			return false
		})
	}

	var maxLengthTo path

	for _, value := range lengthTo {
		if value.value > maxLengthTo.value {
			maxLengthTo = value
		}
	}

	// Log some debug graph stats
	graphStats := graph.Check(g)
	log.Debug().
		Int("size", graphStats.Size).
		Int("multi", graphStats.Multi).
		Int("weighted", graphStats.Weighted).
		Int("loops", graphStats.Loops).
		Int("isolated", graphStats.Isolated).
		Str("graph", g.String()).
		Ints("topSort", topSort).
		Bool("topSortOk", topSortOk).
		Int64("maxPathValue", maxLengthTo.value).
		Ints("maxPath", maxLengthTo.vertexes).
		Msg("Graph stats")

	// Get path petition IDs
	var pathPetitions []model.BookingPetition
	var petitionIds []string

	for _, vertex := range maxLengthTo.vertexes {
		pathPetitions = append(pathPetitions, bookingPetitionGraphMap[vertex])
		petitionIds = append(petitionIds, bookingPetitionGraphMap[vertex].PetitionId)
	}

	// Calculate stats
	stats := u.statsCalculateUseCase.Execute(pathPetitions)

	maximized.PetitionIds = petitionIds
	maximized.TotalProfit = float64(maxLengthTo.value) / 100
	maximized.AvgNight = stats.AvgNight
	maximized.MinNight = stats.MinNight
	maximized.MaxNight = stats.MaxNight

	return maximized
}

type BookingPetitionMap map[int]model.BookingPetition

func (m BookingPetitionMap) FindTargetKeys(fromCheckInDate time.Time) []int {
	var availableObjectives []int

	for key, value := range m {
		if value.CheckIn.After(fromCheckInDate) || value.CheckIn.Equal(fromCheckInDate) {
			availableObjectives = append(availableObjectives, key)
		}
	}

	return availableObjectives
}
