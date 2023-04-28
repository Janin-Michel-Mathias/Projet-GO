package main

import (
	"math/rand"
	"time"
)

type boat struct{
	position [2]int
	length int
	isVertical bool
}

func generateBoats() []boat{
	var boats []boat;
	lengths := [5]int{5,4,3,3,2};
	var newBoat boat;
	rand.Seed(time.Now().UnixNano());

	for i := 0; i < 5; i++{
		newBoat.length = lengths[i];
		newBoat.isVertical = rand.Intn(2) == 0;
		for{
			if(newBoat.isVertical){
				newBoat.position[0] = rand.Intn(10);
				newBoat.position[1] = rand.Intn(10 - newBoat.length);
			} else {
				newBoat.position[0] = rand.Intn(10 - newBoat.length);
				newBoat.position[1] = rand.Intn(10);
			}
			if(validPos(newBoat, boats)){
				break;
			}
		}
		boats = append(boats, newBoat);
	}
	return boats;
}

func validPos(boat boat, boats []boat) bool{
	x := boat.position[0];
	y := boat.position[1];
	overBoat := false;
	for i := 0; i < boat.length; i++ {
		for _, value := range(boats){
			for i := 0; i < value.length; i++ {
				if(value.isVertical){
					if(x == value.position[0] && y == value.position[1] + i){
						overBoat = true;
					}
				} else{
					if(x == value.position[0] + i && y == value.position[1]){
						overBoat = true;
					}
				}
				if(overBoat){
					return false;
				}
			}
			if(boat.isVertical){
				y++;
			}else{
				x++;
			}
		}
	}
	return true;
}