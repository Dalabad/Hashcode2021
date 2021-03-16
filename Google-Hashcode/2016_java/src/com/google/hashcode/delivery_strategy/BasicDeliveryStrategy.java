package com.google.hashcode.delivery_strategy;

import com.google.hashcode.Simulation;
import com.google.hashcode.exporter.OutputExporter;
import com.google.hashcode.models.*;

/**
 * Created by janosch on 2/11/16.
 */
public class BasicDeliveryStrategy implements DeliveryStrategy {

    public int deliver(Simulation simulation, Order order, OutputExporter outputExporter) throws Exception {

        ProductType product;


        while ((product = order.retrieveFirstProductFromBasket()) != null) {

            for (Warehouse warehouse : simulation.getWarehouses()) {
                if (warehouse.isProductInStock(product)) {
                    warehouse.retrieveProductFromStock(product, 1);

                    int turns = 2 + simulation.calculateDuration(order, warehouse);
                    if (turns <= simulation.getRemainingTurns()) {
                        outputExporter.addCommand(new Drone(1, 10000000), new Command("L"), warehouse, product, 1);
                        outputExporter.addCommand(new Drone(1, 10000000), new Command("D"), order, product, 1);
                    }

                    return 2 * turns;

                }
            }

        }

        return 0;
    }

    @Override
    public void deliverAll(Simulation simulation, OutputExporter outputExporter) throws Exception {

        for (Order order : simulation.getOrders()) {
            try {
                if (simulation.getRemainingTurns() >= 0) {
                    int turns = deliver(simulation, order, simulation.getOutputExporter());
                    int remainingTurns = simulation.getRemainingTurns() - turns;
                    simulation.setRemainingTurns(remainingTurns);
                } else {
                    // No more turns left
                    return;
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }


    }
}
