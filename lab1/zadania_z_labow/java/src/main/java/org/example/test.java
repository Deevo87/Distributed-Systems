package org.example;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class test {
    public static boolean canAssignColors(ArrayList<ArrayList<Integer>> adjacencyList) {
        int n = adjacencyList.size();
        int[] colors = new int[n];
        java.util.Arrays.fill(colors, -1); // Initialize colors with -1 indicating uncolored

        for (int i = 0; i < n; i++) {
            if (colors[i] == -1) { // If the current light is uncolored, assign color to it
                if (!dfs(adjacencyList, colors, i, 0)) { // Start DFS from the current light
                    return false; // If DFS fails, return false
                }
            }
        }
        return true;
    }

    public static boolean dfs(ArrayList<ArrayList<Integer>> adjacencyList, int[] colors, int currentLight, int color) {
        if (colors[currentLight] != -1) { // If the current light is already colored
            return colors[currentLight] == color; // Check if its color matches the expected color
        }

        colors[currentLight] = color; // Assign color to the current light

        for (int neighbor : adjacencyList.get(currentLight)) {
            if (!dfs(adjacencyList, colors, neighbor, 1 - color)) { // DFS to neighbors with opposite color
                return false; // If DFS fails, return false
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        ArrayList<ArrayList<Integer>> adjacencyList = new ArrayList<>();

        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            ArrayList<Integer> neighbors = new ArrayList<>();
            String[] tokens = line.split(" ");
            for (String token : tokens) {
                if (!token.isEmpty()) {
                    neighbors.add(Integer.parseInt(token));
                }
            }
            adjacencyList.add(neighbors);
        }

        scanner.close();

        boolean result = canAssignColors(adjacencyList);
        if (result) {
            System.out.println("It is possible to assign colors to every light.");
        } else {
            System.out.println("It is not possible to assign colors to every light.");
        }
    }
}
