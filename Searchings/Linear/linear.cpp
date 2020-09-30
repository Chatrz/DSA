#include <iostream>
#include <vector>

using namespace std;  

bool linear_search(vector <int> numbers, int target) {
    int index;
    for(index = 0; index < numbers.size(); index++)
        if(numbers[index] == target)
            return true;
    return false;
}

vector <int> number_generator() {
    int start = 50, finish = 110, i;
    vector <int> num_list;
    for(i = start; i < finish; i++)
        num_list.push_back(i); 
    return num_list;    
}

int main(){
    int point = 98;
    linear_search(number_generator(), point) == 1 ? cout << "True" : cout << "False" ;   
}
