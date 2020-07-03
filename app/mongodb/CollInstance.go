package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task/app/struct"
)

// InstanceColl : class of InstanceColl
type InstanceColl struct {
	coll *Coll
}

// CollInit ...
func (instanceColl *InstanceColl) CollInit(coll *Coll) {
	instanceColl.coll = coll
}

// GetAllInstance ...
func (instanceColl *InstanceColl) GetAllInstance() []_struct.Instance {
	cur, err := instanceColl.coll.collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	var all = make([]_struct.Instance, 0)

	for cur.Next(context.Background()) {
		var instance _struct.Instance
		if err = cur.Decode(&instance); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, instance)
		}
	}
	cur.Close(context.Background())

	return all
}

// GetInstancesByProjectID ...
func (instanceColl *InstanceColl) GetInstancesByProjectID(projectID primitive.ObjectID) []interface{} {
	cur, err := instanceColl.coll.collection.Find(context.Background(), bson.M{"projectid": projectID})
	if err != nil {
		fmt.Println(err)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	var all = make([]interface{}, 0)

	for cur.Next(context.Background()) {
		var instance _struct.Instance
		if err = cur.Decode(&instance); err != nil {
			fmt.Println(err)
		} else {
			all = append(all, &instance)
		}
	}
	cur.Close(context.Background())
	return all
}

// ProjectHasInstancesRunning ...
func (instanceColl *InstanceColl) ProjectHasInstancesRunning(projectID primitive.ObjectID) bool {
	cur, err := instanceColl.coll.collection.Find(context.Background(), bson.M{"projectid": projectID, "status": "正在运行"})
	if err != nil {
		fmt.Println(err)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.Background())
	return cur.Next(context.Background())
}

// GetInstanceByID ...
func (instanceColl *InstanceColl) GetInstanceByID(instanceID primitive.ObjectID) (_struct.Instance, error) {
	var instance _struct.Instance
	err := instanceColl.coll.collection.FindOne(context.Background(), bson.M{"_id": instanceID}).Decode(&instance)
	return instance, err
}

// AddInstance ...
func (instanceColl *InstanceColl) AddInstance(instance _struct.Instance) error {
	_, err := instanceColl.coll.collection.InsertOne(context.Background(), instance)
	return err
}

// DelInstance ...
func (instanceColl *InstanceColl) DelInstance(instanceIDs []primitive.ObjectID) error {
	for _, instanceID := range instanceIDs {
		_, err := instanceColl.coll.collection.DeleteOne(context.Background(), bson.M{"_id": instanceID})
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateInstance ...
func (instanceColl *InstanceColl) UpdateInstance(instance _struct.Instance) error {
	update := bson.M{"$set": instance}
	updateOpts := options.Update().SetUpsert(true)
	_, err := instanceColl.coll.collection.UpdateOne(context.Background(), bson.M{"_id": instance.InstanceID}, update, updateOpts)
	return err
}
