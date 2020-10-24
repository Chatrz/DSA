/*
the answer to https://www.hackerrank.com/challenges/torque-and-development/problem
using union-find DS which can be found in https://algs4.cs.princeton.edu/15uf/
*/

import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static class WeightedQuickUnion{
        private int[] id;
        private int[]size;
        public WeightedQuickUnion(int n){
            id=new int[n+1];
            size=new int[n+1];
            for(int i=0;i<n+1;i++){
                id[i]=i;
                size[i]=1;
            }
        }
        private int root(int i){
            while(i!=id[i])i=id[i];
            return i;
        }
        public long calCost(int c_lib, int r_road){
            long cost=0;
            for(int i=1;i<id.length;i++){
                if(id[i]==i){
                    //building a library in one city and build minimum num of roads to                       
                    //others
                    cost +=c_lib;
                    cost +=r_road*(size[i]-1);
                }
            }
            return cost;
        }
        public int treeSize(int p){
            return size[p];
        }
        //union
        public void union(int p, int q){
            int pRoot=root(p);
            int qRoot=root(q);
            if(pRoot==qRoot)return;
            if(size[pRoot]<size[qRoot]){id[pRoot]=qRoot;size[qRoot]+=size[pRoot];}
            else {id[qRoot]=pRoot;size[pRoot]+=size[qRoot];}
        }
    }
    static long roadsAndLibraries(int n, int c_lib, int c_road, int[][] cities) {
        /*
        if building library costs less ,then let's build one in all of cities
        */
        if (c_lib<c_road) return (long) n*c_lib;
        else{
            WeightedQuickUnion groups=new WeightedQuickUnion(n);
            for (int i=0;i<cities.length;i++){
                System.out.println(cities[i][0]+cities[i][1]);
                groups.union(cities[i][0],cities[i][1]);
            }
            return groups.calCost(c_lib,c_road);
        }

    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int q = scanner.nextInt();
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int qItr = 0; qItr < q; qItr++) {
            String[] nmC_libC_road = scanner.nextLine().split(" ");

            int n = Integer.parseInt(nmC_libC_road[0]);

            int m = Integer.parseInt(nmC_libC_road[1]);

            int c_lib = Integer.parseInt(nmC_libC_road[2]);

            int c_road = Integer.parseInt(nmC_libC_road[3]);

            int[][] cities = new int[m][2];

            for (int i = 0; i < m; i++) {
                String[] citiesRowItems = scanner.nextLine().split(" ");
                scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

                for (int j = 0; j < 2; j++) {
                    int citiesItem = Integer.parseInt(citiesRowItems[j]);
                    cities[i][j] = citiesItem;
                }
            }

            long result = roadsAndLibraries(n, c_lib, c_road, cities);

            bufferedWriter.write(String.valueOf(result));
            bufferedWriter.newLine();
        }

        bufferedWriter.close();

        scanner.close();
    }
}
