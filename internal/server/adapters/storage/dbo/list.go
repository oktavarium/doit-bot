package dbo

import (
	"github.com/oktavarium/doit-bot/internal/server/domain/planner"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	DbId        primitive.ObjectID `bson:"_id,omitempty"`
	Id          string             `bson:"id,omitempty"`
	OwnerId     string             `bson:"owner_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
}

func FromDomainList(dl *planner.List) List {
	return List{
		Id:          dl.Id(),
		OwnerId:     dl.OwnerId(),
		Name:        dl.Name(),
		Description: dl.Description(),
	}
}

func (l List) ToDomainList() (*planner.List, error) {
	return planner.RestoreListFromDB(
		l.Id,
		l.OwnerId,
		l.Name,
		l.Description,
	)
}

func ListsToDomainLists(lists []List) ([]*planner.List, error) {
	if lists == nil {
		return nil, nil
	}
	result := make([]*planner.List, 0, len(lists))
	for _, l := range lists {
		domainList, err := l.ToDomainList()
		if err != nil {
			return nil, err
		}
		result = append(result, domainList)
	}
	return result, nil
}
