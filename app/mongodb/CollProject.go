package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task/app/struct"
)

// ProjectColl : class of ProjectColl
type ProjectColl struct {
	coll *Coll
}

// AddProject ...
func (projectColl *ProjectColl) AddProject(project _struct.Project) error {
	_, err := projectColl.coll.collection.InsertOne(context.Background(), project)
	return err
}

// DelProject ...
func (projectColl *ProjectColl) DelProject(objectID primitive.ObjectID) error {
	_, err := projectColl.coll.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}

// UpdateProject ...
func (projectColl *ProjectColl) UpdateProject(project _struct.Project) error {
	update := bson.M{"$set": project}
	updateOpts := options.Update().SetUpsert(true)
	_, err := projectColl.coll.collection.UpdateOne(context.Background(), bson.M{"_id": project.ProjectID}, update, updateOpts)
	return err
}

// FindOne ...
func (projectColl *ProjectColl) FindOne(condition primitive.M) (*_struct.Project, error) {
	var project = new(_struct.Project)
	err := projectColl.coll.collection.FindOne(context.Background(), condition).Decode(project)
	if err != nil {
		fmt.Println(err)
		fmt.Println("err: projectColl.go - FindOne")
	}
	return project, err
}

// GetProjectByID ...
func (projectColl *ProjectColl) GetProjectByID(projectID primitive.ObjectID) (_struct.Project, error) {
	var project _struct.Project
	err := projectColl.coll.collection.FindOne(context.Background(), bson.M{"_id": projectID}).Decode(&project)
	return project, err
}

// GetAllInstance ...
func (projectColl *ProjectColl) GetAllProject() []_struct.Project {
	cur, err := projectColl.coll.collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	var all = make([]_struct.Project, 0)

	for cur.Next(context.Background()) {
		var project _struct.Project
		if err = cur.Decode(&project); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, project)
		}
	}
	cur.Close(context.Background())

	return all
}

// CollInit ...
func (projectColl *ProjectColl) CollInit(coll *Coll) {
	projectColl.coll = coll
}
