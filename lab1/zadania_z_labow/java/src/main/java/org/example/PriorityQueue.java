package org.example;

import java.util.ArrayList;

public class PriorityQueue {
    private ArrayList<Integer> heap;

    public PriorityQueue() {
        heap = new ArrayList<>();
    }

    private int parent(int i) {
        return (i - 1) / 2;
    }

    private int left(int i) {
        return 2 * i + 1;
    }

    private int right(int i) {
        return 2 * i + 2;
    }

    private void swap(int i, int j) {
        int tmp = heap.get(i);
        heap.set(i, heap.get(j));
        heap.set(j, tmp);
    }

    private void repairUp(int i) {
        while (i > 0 && heap.get(parent(i)) < heap.get(i)) {
            swap(parent(i), i);
            i = parent(i);
        }
    }

    private void repairDown(int i) {
        int maxInd = i;

        int left = left(i);
        if (left < heap.size() && heap.get(left) > heap.get(maxInd)) {
            maxInd = left;
        }

        int right = right(i);
        if (right < heap.size() && heap.get(right) > heap.get(maxInd)) {
            maxInd = right;
        }

        if (i != maxInd) {
            swap(i, maxInd);
            repairDown(maxInd);
        }
    }

    public void insert(int p) {
        heap.add(p);
        repairUp(heap.size() - 1);
    }

    private int extractMax() {
        int result = heap.get(0);
        heap.set(0, heap.get(heap.size() - 1));
        heap.remove(heap.size() - 1);
        repairDown(0);
        return result;
    }

    public int getMax() {
        return heap.get(0);
    }

    public void remove(int i) {
        heap.set(i, getMax() + 1);
        repairUp(i);
        extractMax();
    }

    public ArrayList<Integer> getHeap() {
        return heap;
    }

    public void printQueue() {
        System.out.println("Priority queue: ");
        for (int e : getHeap()) {
            System.out.println(e);
        }
    }

    public static void main(String[] args) {
        PriorityQueue queue = new PriorityQueue();

        queue.insert(45);
        queue.insert(20);
        queue.insert(14);
        queue.insert(12);
        queue.insert(31);
        queue.insert(7);
        queue.insert(11);
        queue.insert(13);
        queue.insert(7);

        System.out.println("Max priority: " + queue.extractMax());
        queue.printQueue();

    }

}
