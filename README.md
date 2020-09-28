# Parking lot project

## Usage guide
1. Setup project with command below (this command will test, compile, build file in bin)
```
bin/setup
```

2. Parking lot can input 2 types
   - File
     - ```bin/parking_lot ${DIR_FILE}```
     - e.g. ```bin/parking_lot file_inputs.txt```
   - Standard input
     - ```bin/parking_lot```

3. Command to play with
   - ```create_parking_lot ${number}``` for create parking lot size
   - ```park ${registration_number} ${car_colour}``` for park a car at parking lot
   - ```leave ${parking_lot_no}``` for leave a car from parking lot
   - ```status``` for listing only parking lot that was park
   - ```registration_numbers_for_cars_with_colour ${car_colour}``` for listing only registration number that match car color in input
   - ```slot_numbers_for_cars_with_colour ${car_colour}``` for listing only parking lot number that match car color in input
   - ```slot_number_for_registration_number ${registration_number}``` for listing only parking lot number that match registration number in input

4. Run in docker
   - ```docker build . -t ${image_name}```
   - ```docker run -it --name ${container_name} ${image_name}``` with standard input command type
   - ```docker run --name ${container_name} -e CMD=file_inputs.txt ${image_name}``` with file input type