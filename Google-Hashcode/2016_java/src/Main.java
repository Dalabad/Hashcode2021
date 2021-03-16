import com.google.hashcode.Importer.InputImporter;
import com.google.hashcode.Simulation;
import com.google.hashcode.delivery_strategy.DeliveryStrategy;
import com.google.hashcode.exporter.OutputExporter;

import java.io.File;
import java.io.IOException;
/**
 * Created by janosch on 2/11/16.
 */
public class Main {

    public static void main (String[] args) {

        DeliveryStrategy strategy = null;

        // Check if two arguments are gotten
        if (args.length < 2) {
            help();
        }

        // Check if Strategy exists
        String strategyName = args[1];
        try {
            Class strategyClass = Class.forName( "com.google.hashcode.delivery_strategy." + strategyName );
            strategy = (DeliveryStrategy) strategyClass.newInstance();
        } catch (ClassNotFoundException | InstantiationException | IllegalAccessException e) {
            e.printStackTrace();
            help();
        }

        // Check if Input file exitst
        String filename = args[0];
        File inputFile = new File("input", filename);

        if (!inputFile.exists()) {
            help();
        }

        // Import input file
        InputImporter inputImporter = new InputImporter();
        Simulation simulation = inputImporter.readFileData(inputFile);

        // Set strategy
        simulation.setDeliveryStrategy(strategy);

        // Create output writer
        File outputFile = new File("output", filename);
        OutputExporter outputExporter = null;
        try {
            outputExporter = new OutputExporter(outputFile);
        } catch (IOException e) {
            System.out.print("Cannot create output file");
            e.printStackTrace();
            System.exit(-1);
        }

        // Set output exporter for simulation
        simulation.setOutputExporter(outputExporter);

        // Run simulation
        simulation.run();

        // Write output file
        outputExporter.writeOutput();
    }

    public static void help() {
        System.out.println("Please run the following: ./main.jar inputfile strategy");
        System.exit(-1);
    }

}
