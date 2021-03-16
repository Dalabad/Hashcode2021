package com.google.hashcode.models;

import java.util.*;

/**
 * Created by neotreat on 11/02/16.
 */
public class Order extends Location implements Cloneable {

    /**
     *
     */
    private int id;

    /**
     *
     */
    public Order(int id) {
        this.id = id;
    }

    /**
     *
     */
    private HashMap<ProductType, Integer> basket = new HashMap<ProductType, Integer>();

    /**
     * @return
     */
    public int getId() {
        return id;
    }

    /**
     * @param id
     */
    public void setId(int id) {
        this.id = id;
    }

    /**
     *
     */
    public void setBasket(HashMap<ProductType, Integer> basket) {
        this.basket = basket;
    }

    /**
     *
     */
    public void addBasket(HashMap<ProductType, Integer> basket) {

        // Add all products per type
        for (Map.Entry<ProductType, Integer> e : basket.entrySet()) {

            ProductType product = e.getKey();
            Integer amount = e.getValue();

            this.addProductToBasket(product, amount);
        }
    }

    /**
     * @param product
     */
    public void addProductToBasket(ProductType product, Integer amount) {

        // Check if the product type is already present in the payload
        if (this.basket.containsKey(product)) {
            Integer currentAmount = this.basket.get(product);
            this.basket.put(product, new Integer(currentAmount.intValue() + amount.intValue()));
        } else {
            this.basket.put(product, amount);
        }
    }

    /**
     *
     * @return
     */
    public HashMap<ProductType, Integer> retrieveMaxPayload(int weight) {

        HashMap<ProductType, Integer> retrievedProducts = new HashMap<ProductType, Integer>();
        int remainingWeight = weight;

        // Get all available products
        SortedSet<ProductType> products = new TreeSet<ProductType>(this.basket.keySet());

        for(final Iterator it = products.iterator(); it.hasNext(); )
        {
            ProductType product = (ProductType) it.next();

            int amount = 1;

            if(product.getWeight() <= remainingWeight) {

                remainingWeight -= product.getWeight();

                // Check how many of these items fit into the drone
                // TODO: Check if there are more items available in basket, only then proceed to retrieve more
                while(remainingWeight >= product.getWeight()) {
                    // remove more items
                    amount++;
                    remainingWeight -= product.getWeight();
                }

                try {
                    this.retrieveProductFromBasket(product, amount);
                    retrievedProducts.put(product, amount);
                } catch (Exception e) {
                    e.printStackTrace();
                }

                return retrievedProducts;
            }
        }

        // Iterate over all product types


        return null;
    }

    /**
     *
     * @return
     */
    public HashMap<ProductType, Integer> getMaxPayload(int weight) {

        HashMap<ProductType, Integer> retrievedProducts = new HashMap<ProductType, Integer>();
        int remainingWeight = weight;

        // Get all available products
        SortedSet<ProductType> products = new TreeSet<ProductType>(this.basket.keySet());

        for(final Iterator it = products.iterator(); it.hasNext(); )
        {
            ProductType product = (ProductType) it.next();

            int amount = 1;

            if(product.getWeight() <= remainingWeight) {

                remainingWeight -= product.getWeight();

                // Check how many of these items fit into the drone
                while(remainingWeight >= product.getWeight()) {
                    // remove more items
                    amount++;
                    remainingWeight -= product.getWeight();
                }

                retrievedProducts.put(product, new Integer(amount));

                return retrievedProducts;
            }
        }

        // Iterate over all product types


        return null;
    }

    /**
     * Check if a product is in the basket.
     */
    public boolean isProductInBasket(ProductType product) {
        return PayloadHelper.isProductInMap(product, this.basket);
    }

    /**
     * Take a product from the basket.
     */
    public void retrieveProductFromBasket(ProductType product, Integer amount) throws Exception {

        // Check if the item is in stock
        if (!isProductInBasket(product)) {
            throw new Exception("Item out of stock!");
        }

        // Current number of products
        Integer currentAmount = this.basket.get(product);

        // Check if there are that many items available
        if(currentAmount.intValue() < amount.intValue()) {
            throw new Exception("Not that many items available");
        }

        // If there is nothing left remove the product completely
        Integer newAmount = new Integer(currentAmount.intValue() - amount.intValue());

        if (newAmount.intValue() <= 0) {
            this.basket.remove(product);
        }
        else {
            this.basket.put(product, newAmount);
        }
    }

    /**
     *
     *
     * @return
     */
    public ProductType getFirstProduct() {

        // Get all available products
        Object[] products = this.basket.keySet().toArray();

        // No more products are available
        if (products.length == 0) {
            return null;
        }
        else {
            return (ProductType) products[0];
        }

    }

    /**
     *
     */
    public ProductType retrieveFirstProductFromBasket() {

        // Get all available products
        Object[] products = this.basket.keySet().toArray();

        // No more products are available
        if (products.length == 0) {
            return null;
        }
        else {

            // Get the first product
            ProductType firstProduct = (ProductType) products[0];

            try {
                this.retrieveProductFromBasket(firstProduct, new Integer(1));
                return firstProduct;
            } catch (Exception e) {
                e.printStackTrace();
                return null;
            }
        }
    }

    public Object clone() throws CloneNotSupportedException {

        Order cloneOrder = (Order) super.clone();
        cloneOrder.setBasket((HashMap<ProductType, Integer>) basket.clone());

        return cloneOrder;
    }
}
