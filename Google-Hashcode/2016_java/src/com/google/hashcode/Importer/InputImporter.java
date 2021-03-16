/**
 * Created by Daniel Schosser on 2/11/16.
 */
package com.google.hashcode.Importer;

import com.google.hashcode.Simulation;
import com.google.hashcode.models.*;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class InputImporter {

    private int maxTurns;
    private int amountOfDrones;
    private int maxPayload;
    
    private boolean logging = false;

    public Simulation readFileData(File file) {

        int lineNumber = 0;
        int amountOfWarehouses = -1;
        int amountOfOrders = -1;
        int amountOfProducts = -1;

        List<Warehouse> warehouses = new ArrayList<Warehouse>();
        List<Order> orders = new ArrayList<Order>();
        List<ProductType> productTypes = new ArrayList<ProductType>();
        List<Drone> drones = new ArrayList<Drone>();

        boolean warehousePositionIsSet = false;
        Warehouse tmpWarehouse = null;
        Order tmpOrder = null;
        int tmpOrderItemCount = -1;

        try (BufferedReader br = new BufferedReader(new FileReader(file))) {
            String line;

            // read line by line
            while ((line = br.readLine()) != null) {

                switch (lineNumber) {
                    case 0:
                        log("Rows, columns, ... " + line);
                        extractSettings(line);
                        break;
                    case 1:
                        log("Amount of Products " + line);
                        amountOfProducts = Integer.parseInt(line);
                        break;
                    case 2:
                        log("Producttypes " + line);
                        productTypes = extractProductWeights(amountOfProducts, line);
                        break;
                    case 3:
                        log("Amount of Warehouses " + line);
                        amountOfWarehouses = Integer.parseInt(line);
                        break;
                    default:
                        String[] splittedLine = line.split(" ");

                        // read warehouse data
                        if(amountOfWarehouses>0) {

                            // create new warehouse and set its location
                            if(tmpWarehouse == null) {
                                log("Create Warehouse #" + warehouses.size() + " and set its position: " + line);
                                tmpWarehouse = new Warehouse(warehouses.size());
                                tmpWarehouse.row = Integer.parseInt(splittedLine[0]);
                                tmpWarehouse.column = Integer.parseInt(splittedLine[1]);
                            } else { // set warehouse stock
                                log("Add items to Warehouse #" + warehouses.size() + ": " + line);
                                for (int i = 0; i < productTypes.size(); i++) {
                                    tmpWarehouse.addProductToStock(productTypes.get(i), Integer.parseInt(splittedLine[i]));
                                }

                                // add warehouse to list
                                warehouses.add(tmpWarehouse);

                                log("Reset tmpWarehouse data");
                                // reset tmpWarehouse
                                amountOfWarehouses--;
                                tmpWarehouse = null;
                            }
                            break;
                        }

                        // set amount of orders
                        if(amountOfOrders == -1) {
                            log("Set amount of Orders: " + line);
                            amountOfOrders = Integer.parseInt(line);
                            break;
                        }

                        // read order data
                        if(amountOfOrders>0) {

                            // create new order and set its destination
                            if(tmpOrder == null) {
                                log("Create order #" + orders.size() + ": " + line);
                                tmpOrder = new Order(orders.size());
                                tmpOrder.row = Integer.parseInt(splittedLine[0]);
                                tmpOrder.column = Integer.parseInt(splittedLine[1]);
                            } else if(tmpOrderItemCount == -1) { // read item count for order
                                log("Set amount of items for order #" + orders.size() + ": " + line);
                                tmpOrderItemCount = Integer.parseInt(line);
                            } else {
                                log("Add items to order #" + orders.size() + ": " + line);
                                // add products to order
                                for (String s : splittedLine) {
                                    tmpOrder.addProductToBasket(productTypes.get(Integer.parseInt(s)), 1);
                                }

                                // add order to list
                                orders.add(tmpOrder);

                                // reset tmpOrder
                                log("Reset tmpOrder");
                                amountOfOrders--;
                                tmpOrder = null;
                                tmpOrderItemCount = -1;
                            }
                            break;
                        }
                        break;
                }

                lineNumber++;
            }

            for(int i=0;i<amountOfDrones;i++) {
                Drone drone = new Drone(i, maxPayload);
                drone.row = warehouses.get(0).row;
                drone.column = warehouses.get(0).column;
                drone.setFlightState(DroneFlightStates.HOVERING_WAREHOUSE);
                drones.add(drone);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }

        return new Simulation(maxTurns, drones, warehouses, orders, productTypes);
    }

    private List<ProductType> extractProductWeights(int amount, String line) {
        List<ProductType> productTypes = new ArrayList<ProductType>();
        String[] weightsArray = line.split(" ");

        for (int i = 0; i < amount; i++) {
            productTypes.add(new ProductType(i, Integer.parseInt(weightsArray[i])));
        }

        return productTypes;
    }

    private void extractProductTypes(String line) {

    }

    public void extractSettings(String line) {
        String[] settingsArray = line.split(" ");

        // int rows = Integer.parseInt(settingsArray[0]);
        // int columns = Integer.parseInt(settingsArray[1]);
        amountOfDrones = Integer.parseInt(settingsArray[2]);
        maxTurns = Integer.parseInt(settingsArray[3]);
        maxPayload = Integer.parseInt(settingsArray[4]);
    }
    
    private void log(String string) {
        if(logging == true)
            System.out.println(string);
    }
}