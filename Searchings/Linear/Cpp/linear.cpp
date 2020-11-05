#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include <vector>
#include <ctime>

using namespace std;  

const int START = 100;
const int STOP = 200;
const int NUMBERS = 90;
const int TARGET = (rand() % (STOP - START)) + START;
int searches = 0;

bool linear_search(vector <int> numbers, int target) {
    int index;
    for(index = 0; index < numbers.size(); index++, searches++) 
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

int main(){
    srand((unsigned) time(0));
    cout << "Searching for " << TARGET << " in array of " << NUMBERS << " numbers was : ";
    linear_search(number_generator(), TARGET) == 1 ? cout << "True" : cout << "False" ;
    cout << "\nTotal searches -> " << searches << endl;
}
