package com.google.hashcode.delivery_strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.models.*;

import java.util.List;

/**
 * Created by janosch on 2/11/16.
 */
public class MultiDroneDeliveryStrategy implements DeliveryStrategy {

    private List<Drone> drones;

    private int calculateWay(Simulation simulation, Order order, Drone drone) {

        ProductType product;
        int turns = 0;
        Order dummyOrder;
        Drone dummyDrone;


        try {
            dummyOrder = (Order) order.clone();
            dummyDrone = (Drone) drone.clone();
        } catch (CloneNotSupportedException e) {
            e.printStackTrace();
            return -1;
        }

        while ((product = dummyOrder.retrieveFirstProductFromBasket()) != null) {

            for (Warehouse warehouse : simulation.getWarehouses()) {
                if (warehouse.isProductInStock(product)) {
                    turns += simulation.calculateDuration(dummyDrone, warehouse) + 2 + simulation.calculateDuration(order, warehouse);
                    dummyDrone.row = order.row;
                    dummyDrone.column = order.column;
                    break;
                }
            }

        }

        return turns;

    }

    public int deliver(Simulation simulation, Order order, Drone drone, OutputExporter outputExporter) throws Exception {

        ProductType product;

        while ((product = order.retrieveFirstProductFromBasket()) != null) {

            for (Warehouse warehouse : simulation.getWarehouses()) {
                if (warehouse.isProductInStock(product)) {
                    warehouse.retrieveProductFromStock(product, 1);

                    int turns = simulation.calculateDuration(drone, warehouse) + 2 + simulation.calculateDuration(order, warehouse);

                    outputExporter.addCommand(drone, new Command("L"), warehouse, product, 1);
                    outputExporter.addCommand(drone, new Command("D"), order, product, 1);

                    drone.row = order.row;
                    drone.column = order.column;

                    drone.addTurns(turns);

                    break;
                }

            }

        }

        return 0;
    }

    @Override
    public void deliverAll(Simulation simulation, OutputExporter outputExporter) throws Exception {

        int numberDrones = simulation.getDrones().size() - 1;

        int i = 0;

        for (Order order : simulation.getOrders()) {
            try {

                int thisDrone = i;

                Drone drone = simulation.getDrones().get(i);
                int turns = calculateWay(simulation, order, drone);

                while(drone.getTurn() + turns > simulation.getMaxTurns()) {
                    i += 1;
                    i = i % numberDrones;
                    if (i == thisDrone) {
                        // No more time to deliver
                        return;
                    }
                    drone = simulation.getDrones().get(i);
                    turns = calculateWay(simulation, order, drone);
                }

                deliver(simulation, order, drone, simulation.getOutputExporter());

            } catch (Exception e) {
                e.printStackTrace();
            }
        }


    }


}
