package main

import (
	"log"
	"regexp"
	_ "regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarPlateGenerator(t *testing.T) {
	generator := NewCarPlateGenerator()

	for i := 0; i < 10; i++ {
		plate := generator.Generate()
		assert.Regexp(t, "^[A-Z]{3}-\\d{4}$", plate, "Invalid car plate format")
	}
}

func TestCarFactory(t *testing.T) {
	car := CarFactory()

	assert.NotEmpty(t, car.CarPlate, "Car plate should not be empty")
	assert.NotEmpty(t, car.CarType, "Car type should not be empty")
	assert.NotEmpty(t, car.CarColor, "Car color should not be empty")
	assert.NotEmpty(t, car.CarBrand, "Car brand should not be empty")
	assert.NotEmpty(t, car.VehicleOwnerCPF, "Vehicle owner CPF should not be empty")

}

func TestTrafficGenerator(t *testing.T) {
	trafficGen := NewTrafficGenerator()

	assert.NotNil(t, trafficGen.TrafficInfo.CarPlate, "Car plate should not be nil")
	assert.NotNil(t, trafficGen.TrafficInfo.CarType, "Car type should not be nil")
	assert.NotNil(t, trafficGen.TrafficInfo.CarColor, "Car color should not be nil")
	assert.NotNil(t, trafficGen.TrafficInfo.CarBrand, "Car brand should not be nil")
	assert.NotNil(t, trafficGen.TrafficInfo.VehicleOwnerCPF, "Vehicle owner CPF should not be nil")


}


func TestDateGenerator(t *testing.T) {
	generator := NewDateGenerator()

	
	for i := 0; i < 10; i++ {
		date := generator.Generate()
		log.Print("Date: ", date)

		
		match := regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}$`).MatchString(date)
		assert.True(t, match, "Invalid date format: %s", date)
	}
}

func TestMaxSpeedGenerator(t *testing.T) {
	generator := NewMaxSpeedGenerator()

	
	for i := 0; i < 10; i++ {
		maxSpeed := generator.Generate()
		assert.GreaterOrEqual(t, maxSpeed, 30, "Max speed should be greater than or equal to 30")
		assert.LessOrEqual(t, maxSpeed, 210, "Max speed should be less than or equal to 210")
	}
}

func TestTimeGenerator(t *testing.T) {
	generator := NewTimeGenerator()

	

		time := generator.Generate()
		log.Println("Time: ", time)

		
		match := regexp.MustCompile(`^\d{1,2}:\d{2}:\d{2}$`).MatchString(time)
		assert.True(t, match, "Invalid time format: %s", time)

}


func TestSpeedGenerator(t *testing.T) {
	generator := NewSpeedGenerator()

	if(generator.Generate() > 0){
		assert.NotEmpty(t, generator.Generate(), "Speed should not be empty")
	}
	lowerThan210 := generator.Generate() < 210
	log.Println("Speed: ", generator.Generate())
	assert.True(t, lowerThan210, "Speed should be lower than 210")

}

func TestAddressGenerator(t *testing.T) {
	generator := NewAddressGenerator()

	
	for i := 0; i < 10; i++ {
		address := generator.Generate()
		assert.NotEmpty(t, address, "Address should not be empty")
	}
}

func TestDirectionGenerator(t *testing.T) {
	generator := NewDirectionGenerator()

	
	for i := 0; i < 10; i++ {
		direction := generator.Generate()
		assert.Contains(t, []string{"North", "South", "East", "West"}, direction, "Invalid direction")
	}
}

func TestStreetDirectionGenerator(t *testing.T) {
	generator := NewStreetDirectionGenerator()

	
	for i := 0; i < 10; i++ {
		streetDirection := generator.Generate()
		assert.Contains(t, []string{"North", "South", "East", "West"}, streetDirection, "Invalid street direction")
	}
}