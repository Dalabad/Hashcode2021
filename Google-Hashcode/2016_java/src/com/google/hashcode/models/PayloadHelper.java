package com.google.hashcode.models;

import java.util.HashMap;
import java.util.Map;

/**
 * Created by neotreat on 11/02/16.
 */
public class PayloadHelper {

    /**
     *
     * @param payload
     * @return
     */
    public static int calculatePayloadWeight (HashMap<ProductType, Integer> payload) {

        // Weight of all payload products
        int weight = 0;

        for(Map.Entry<ProductType, Integer> e: payload.entrySet()){
            ProductType product = e.getKey();
            Integer amount = e.getValue();

            weight += product.getWeight() * amount.intValue();
        }

        return weight;
    }

    /**
     * Check if a product is in stock.
     */
    public static boolean isProductInMap(ProductType product, HashMap<ProductType, Integer> map) {

        // Check if this product was intialised
        if (!map.containsKey(product))
        {
            return false;
        }

        // Amount of items in stock for this product type
        Integer amount = map.get(product);

        if(amount.intValue() > 0) {
            return true;
        }
        else
        {
            return false;
        }
    }
}
