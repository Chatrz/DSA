#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <vector>
#include <ctime>
#include <thread>

using namespace std;  

const int START = 100;
const int STOP = 200;
const int NUMBERS = 90;
const int THREADS = 10;

bool linear_search(vector <int> numbers, int target, int &search) {
    int index;
    for(index = 0; index < numbers.size(); index++, search++) 
        if(numbers[index] == target)
            return true;
    return false;
}

vector <int> number_generator() {
    vector <int> num_list;
    for(int i = 0; i < NUMBERS; i++) 
        num_list.push_back((rand() % (STOP - START)) + START); 
    return num_list;    
}

void LinearSearch(int threadid, int target) {
    int tid = threadid;
    int searches = 0;
    cout << "<< Thread " << tid << " ,started searching for " << target << " in array of " << NUMBERS << " numbers." << endl;
    string result = linear_search(number_generator(), target, searches) == 1 ? "True" : "False";
    cout << ">> Thread " << tid << " ,finished searching with result = " << result << " / Total searches -> " << searches << "\n" <<endl;
}

int main(){
    srand((unsigned) time(0));
    for( int i = 0; i < THREADS; i++) {
        thread t(LinearSearch, i+100, (rand() % (STOP - START)) + START);
        t.join();
    }
}
