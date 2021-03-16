package com.google.hashcode.models;

import java.util.HashMap;
import java.util.Map;

/**
 * Created by neotreat on 11/02/16.
 */
public class Drone extends Location implements Cloneable {

    /**
     *
     */
    private int id;

    /**
     * Maximum weight of all payloads.
     */
    private int maxPayloadWeight;

    /**
     *
     */
    private HashMap<ProductType, Integer> payload = new HashMap<ProductType, Integer>();

    /**
     *
     */
    private int flightState = DroneFlightStates.HOVERING_WAREHOUSE;

    /*
     *
     */
    private int turn = 0;

    /**
     *
     */
    public Drone(int id, int maxPayloadWeight) {
        this.id = id;
        this.maxPayloadWeight = maxPayloadWeight;
    }

    /**
     *
     * @return
     */
    public int getId() {
        return id;
    }

    /**
     *
     * @param id
     */
    public void setId(int id) {
        this.id = id;
    }

    /**
     *
     * @return
     */
    public int getMaxPayloadWeight() {
        return maxPayloadWeight;
    }

    /**
     *
     * @param maxPayloadWeight
     */
    public void setMaxPayloadWeight(int maxPayloadWeight) {
        this.maxPayloadWeight = maxPayloadWeight;
    }

    /**
     * Set the payload, which will replace the previous one completely!
     */
    public void setPayload(HashMap<ProductType, Integer> payload) throws Exception {

        // Check if the payload is too heavy
        if (PayloadHelper.calculatePayloadWeight(payload) > this.maxPayloadWeight) {
            throw new Exception("The payload is too heavy for this drone.");
        }

        this.payload = payload;
    }

    /**
     *
     * @return
     */
    public int getFlightState() {
        return flightState;
    }

    /**
     *
     * @param flightState
     */
    public void setFlightState(int flightState) {
        this.flightState = flightState;
    }

    /**
     *
     * @param payload
     * @throws Exception
     */
    public void addPayload(HashMap<ProductType, Integer> payload) throws Exception {

        // Check if the overall payload is too heavy
        if(PayloadHelper.calculatePayloadWeight(payload) + this.getPayloadWeight() > this.maxPayloadWeight) {
            throw new Exception("The payload is too heavy for this drone.");
        }

        // Add all products per type
        for(Map.Entry<ProductType, Integer> e: payload.entrySet()){

            ProductType product = e.getKey();
            Integer amount = e.getValue();

            this.addProductToPayload(product, amount);
        }
    }

    /**
     *
     * @param product
     */
    public void addProductToPayload(ProductType product, Integer amount) throws Exception {

        // Check if the overall payload is too heavy
        if(product.getWeight() * amount.intValue() + this.getPayloadWeight() > this.maxPayloadWeight) {
            throw new Exception("The payload is too heavy for this drone.");
        }

        // Check if the product type is already present in the payload
        if(this.payload.containsKey(product)) {
            Integer currentAmount = this.payload.get(product);
            this.payload.put(product, new Integer(currentAmount.intValue() + amount.intValue()));
        }
        else {
            this.payload.put(product, amount);
        }
    }

    /**
     * Check if a product is in the payload.
     */
    public boolean isProductInPayload(ProductType product) {
        return PayloadHelper.isProductInMap(product, this.payload);
    }

    /**
     * Take a product from the payload.
     */
    public void retrieveProductFromPayload(ProductType product, Integer amount) throws Exception {

        // Check if the item is in stock
        if (!isProductInPayload(product)) {
            throw new Exception("Item out of stock!");
        }

        // Current number of products
        Integer currentAmount = this.payload.get(product);

        // Check if there are that many items available
        if(currentAmount.intValue() < amount.intValue()) {
            throw new Exception("Not that many items available");
        }

        // If there is nothing left remove the product completely
        Integer newAmount = new Integer(currentAmount.intValue() - amount.intValue());

        if (newAmount.intValue() <= 0) {
            this.payload.remove(product);
        }
        else {
            this.payload.put(product, newAmount);
        }
    }

    /**
     *
     */
    public int getPayloadWeight() {
        return PayloadHelper.calculatePayloadWeight(this.payload);
    }

    public int getTurn() {
        return turn;
    }

    public void addTurns(int turns) {
        this.turn += turns;
    }

    public void decreaseTurns() { this.turn--; }

    public Object clone() throws CloneNotSupportedException {

        Drone cloneDrone = (Drone) super.clone();
        try {
            cloneDrone.setPayload(this.payload);
        } catch (Exception e) {
            e.printStackTrace();
        }

        return super.clone();
    }

    // TODO: If drone is at warehouse, at items to warehouse after dropping them from the drone
    public void dropAllItems() {
        payload = new HashMap<ProductType, Integer>();
    }
}
