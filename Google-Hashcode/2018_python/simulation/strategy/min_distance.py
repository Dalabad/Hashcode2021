class MinDistance:

    def run(self, city):
        city.rides = sorted(city.rides, key=lambda ride: ride.earlieststart)

        for ride in city.rides:

            min_distance = city.rows * 2
            car_id = city.cars

            for car in city.cars:

                distance_car_to_ride = car.get_distance(ride)
                
                if distance_car_to_ride < min_distance and car.get_ride_end_time(ride) <= ride.latestend:
                    min_distance = distance_car_to_ride
                    car_id = car.id

            if car_id == city.cars:
                continue

            city.cars[car_id].rides.append(ride)
            city.cars[car_id].simulation_step = city.cars[car_id].get_ride_end_time(ride)
            city.cars[car_id].position = ride.end

        return city
