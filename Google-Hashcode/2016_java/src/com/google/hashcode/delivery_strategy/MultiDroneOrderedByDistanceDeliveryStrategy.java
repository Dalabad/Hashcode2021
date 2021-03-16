package com.google.hashcode.delivery_strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.models.*;

import java.util.List;
import java.util.Map;

/**
 * Created by daniel on 2/11/16.
 */
public class MultiDroneOrderedByDistanceDeliveryStrategy implements DeliveryStrategy {

    public int deliver(Simulation simulation, Order order, OutputExporter outputExporter) throws Exception {
        return 0;
    }

    @Override
    public void deliverAll(Simulation simulation, OutputExporter outputExporter) throws Exception {

        for(int i = 0; i < simulation.getMaxTurns(); i++) {
            List<Order> orders = simulation.getOrders();
            List<Drone> drones = simulation.getDrones();

            for(Drone drone : drones) {

                if(drone.getFlightState() == DroneFlightStates.HOVERING_CUSTOMER) {
                    Warehouse warehouse = simulation.findNearestWarehouse(drone);
                    drone.setFlightState(DroneFlightStates.FLYING);
                    drone.addTurns(simulation.calculateDuration(drone, warehouse));

                    Order order = findNearestOrder(simulation, orders, drone);

                    int freePayload = drone.getMaxPayloadWeight() - drone.getPayloadWeight();
                    for(Map.Entry<ProductType, Integer> e: order.getMaxPayload(freePayload).entrySet()){
                        ProductType product = e.getKey();
                        Integer amount = e.getValue();

                        drone.addProductToPayload(product, amount);
                        outputExporter.addCommand(drone, new Command("L"), order, product, amount);
                    }

                } else if(drone.getFlightState() == DroneFlightStates.HOVERING_WAREHOUSE) {

                    Order order = findNearestOrder(simulation, orders, drone);

                    if(order != null) {
                        if(simulation.getRemainingTurns() > simulation.calculateDuration( drone, order )) {
                            int freePayload = drone.getMaxPayloadWeight() - drone.getPayloadWeight();
                            int orderIndex = orders.indexOf(order);

                            for(Map.Entry<ProductType, Integer> e: order.retrieveMaxPayload(freePayload).entrySet()){
                                ProductType product = e.getKey();
                                Integer amount = e.getValue();

                                drone.addProductToPayload(product, amount);
                                outputExporter.addCommand(drone, new Command("D"), order, product, amount);
                            }

                            orders.set(orderIndex, order);
                            drone.setFlightState(DroneFlightStates.FLYING);
                            drone.addTurns(simulation.calculateDuration( drone, order));
                        }
                    }
                } else {
                    drone.decreaseTurns();

                    if(drone.getTurn() == 0) {
                        if(simulation.isAtWarehouse(drone)) {
                            drone.setFlightState(DroneFlightStates.HOVERING_WAREHOUSE);
                        } else {
                            drone.setFlightState(DroneFlightStates.HOVERING_CUSTOMER);
                            // drop all items from drone, as it only contains items to the current customer
                            drone.dropAllItems();
                        }
                    }
                }
            }

            simulation.setRemainingTurns(simulation.getRemainingTurns()-1);
        }
    }


    public Order findNearestOrder(Simulation simulation, List<Order> orders, Location location) {
        Order nearest = null;
        int nearestDuration = Integer.MAX_VALUE;
        for(Order order : orders) {
            int duration = simulation.calculateDuration(location, order);
            if(duration < nearestDuration) {
                nearest = order;
                nearestDuration = duration;
            }
        }

        return nearest;
    }

}
