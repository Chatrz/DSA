#include <iostream>
#include <queue>
using namespace std;

int main()
{
  queue<int> item;
  queue<int> count;
  int x , y,n ;
  char c;
  cin >> n;
  for (int i=0;i<n;i++) {
    cin >> c;
    if (c=='+'){
      cin >> x;
      cin >> y;
      item.push(x);
      count.push(y);
    }else if (c =='-'){
      cin >> x;
      while (!count.empty() && x!=0)
      {
          count.front() -= x;
          if (count.front()<=0){
              x = -count.front();
              count.pop();
              item.pop();
          }
          else
              x = 0;
      }
    }else{
      if (!count.empty())
        cout << item.front() << "\n";
      else
        cout << "empty\n";
    }
  }
}
