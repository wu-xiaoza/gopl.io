package main

import (
	"ORM-net/ent"
	"ORM-net/ent/car"
	"ORM-net/ent/group"
	"ORM-net/ent/user"
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:1990686992@/dbtest?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	//Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	ctx := context.Background()
	//Run the auto migration tool
	//Schema is the client for creating, migrating and dropping schema.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if _, err := CreateUser(ctx, client); err != nil {
		log.Fatal(err)
	}

	if _, err := QueryUser(ctx, client); err != nil {
		log.Fatal(err)
	}

	a8m, err := CreateCars(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(a8m)

	if err := QueryCars(ctx, a8m); err != nil {
		log.Fatal(err)
	}

	if err = QueryCarUsers(ctx, a8m); err != nil {
		log.Fatal(err)
	}

	if err := CreateGraph(ctx, client); err != nil {
		log.Fatal(err)
	}

	if err := QueryArielCars(ctx, client); err != nil {
		log.Fatal(err)
	}
}

//CreateUser .
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

//QueryUser .
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		//NameEQ applies the EQ predicate on the "name" field.
		Where(user.NameEQ("a8m")).
		//Only returns a single User entity found by the query, ensuring it only returns one. Returns a *NotSingularError when exactly one User entity is not found. Returns a *NotFoundError when no User entities are found.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

//CreateCars .
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	//// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

//QueryCars .
func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.ModelEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(ford)
	return nil
}

//QueryCarUsers .
func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	// Query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("fauled querying car %q owner: %v", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}

//CreateGraph .
func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.Create().SetAge(30).SetName("Ariel").Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.Create().SetAge(28).SetName("Neta").Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).              // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

//QueryGithub .
func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %v", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

//QueryArielCars .
func QueryArielCars(ctx context.Context, client *ent.Client) error {
	// Get "Ariel" from previous steps.
	a8m := client.User.
		Query().
		Where(
			//HasCars applies the HasEdge predicate on the "cars" edge.
			user.HasCars(),
			user.Name("Ariel"),
		).
		//OnlyX is like Only, but panics if an error occurs.
		OnlyX(ctx)
	cars, err := a8m. // Get the groups, that a8m is connected to:
				QueryGroups(). // (Group(Name=GitHub), Group(Name=GitLab),)
				QueryUsers().  // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
				QueryCars().   //
				Where(         //
			//Not applies the not operator on the given predicate.
			car.Not( //  Get Neta and Ariel cars, but filter out
				car.ModelEQ("Mazda"), //  those who named "Mazda"
			), //
		). //
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %v", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
	return nil
}

//QueryGroupWithUsers .
func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %v", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}
