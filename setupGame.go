package main

import (
	"math/rand"
	"time"
)

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
	for i := 0; i < boat.length; i++ {

			if(isBoatPos([2]int{x,y}, boats)){
				return false;
			}	

			if(boat.isVertical){
				y++;
			}else{
				x++;
			}
		}
		return true;
	}
	


func isBoatPos(position [2]int, boats []boat) bool{
	for i:= 0; i < len(boats); i++ {
		x := boats[i].position[0];
		y := boats[i].position[1];
		for j:=0; j < boats[i].length; j++{
			if(position[0] == x && position[1] == y){
				return true
			}
			if(boats[i].isVertical){
				y++;
			}else {
				x++;
			}
		}
	}
	return false;
}