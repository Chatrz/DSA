#include <iostream>
#include <stack>
#include <map>
#include <sstream>
#include <string>
#include <math.h>

using namespace std;

class Calculator
{
public:
    int MAX;
    Calculator(string in);
    void moveCursor(char c);
    void insert(char c);
    void remove();
    long long int evaluate();
    string print();

private:
    struct node
    {
        node() : next(NULL), prev(NULL), C(' '){};
        char C;
        node *next;
        node *prev;
    };
    map<string, int> results;
    node *head;
    node *tail;
    node *cursor;
    void append(char c);
    string getListStr();
    bool is_number(const string &s);
    bool is_operator(char c);
    int prec(char operand);
    string postfix(string exp);
    long long int mod(long long int n, bool negetive);
    int longMod(string num);
};

Calculator::Calculator(string in)
{
    MAX = pow(10, 9) + 7;
    head = new node();
    tail = head;
    cursor = head;
    cursor->C = '|';
    for (char c : in)
        append(c);
}

void Calculator::moveCursor(char c)
{
    if (c == '<')
    {
        if (cursor == head)
            return;
        cursor->C = cursor->prev->C;
        cursor->prev->C = '|';
        cursor = cursor->prev;
        return;
    }
    if (cursor->next == NULL)
        return;
    cursor->C = cursor->next->C;
    cursor->next->C = '|';
    cursor = cursor->next;
    return;
}

void Calculator::insert(char c)
{
    cursor->C = c;
    node *temp = new node();
    temp->C = '|';
    temp->prev = cursor;
    temp->next = cursor->next;
    if (cursor->next != NULL)
        cursor->next->prev = temp;
    cursor->next = temp;
    cursor = temp;
    return;
}

void Calculator::remove()
{
    if (cursor->prev == NULL)
        return;
    cursor->prev->C = '|';
    cursor->prev->next = cursor->next;
    if (cursor->next != NULL)
        cursor->next->prev = cursor->prev;
    cursor = cursor->prev;
    return;
}

long long int Calculator::mod(long long int n, bool negetive)
{
    if (negetive && n < 0)
    {
        while (n < 0)
            n += MAX;
        return n;
    }
    return n % MAX;
}

int Calculator::longMod(string num)
{
    int res = 0;
    for (int i = 0; i < num.length(); i++)
        if (num[i] != '+' && num[i] != '-')
            res = (res * 10 + (int)num[i] - '0') % MAX;
    return res;
}

long long int Calculator::evaluate()
{
    string exp = postfix(getListStr());
    if (results.count(exp) > 0)
        return results[exp];
    stack<string> opStack;
    stringstream tokenlist(exp);
    string token;
    while (getline(tokenlist, token, ' '))
    {
        if (is_number(token))
            opStack.push(token);
        else if (prec(token[0]))
        {
            long long int val1 = stoll(opStack.top(), nullptr, 10);
            opStack.pop();
            long long int val2 = stoll(opStack.top(), nullptr, 10);
            opStack.pop();
            switch (token[0])
            {
            case '+':
                opStack.push(to_string(mod(val2 + val1, false)));
                break;
            case '-':
                opStack.push(to_string(mod(val2 - val1, false)));
                break;
            case '*':
                opStack.push(to_string(mod(val2 * val1, false)));
                break;
            }
        }
    }
    results[exp] = mod(stoll(opStack.top(), nullptr, 10), true);
    return mod(stoll(opStack.top(), nullptr, 10), true);
}

void Calculator::append(char c)
{
    node *temp = new node();
    temp->C = c;
    temp->prev = tail;
    tail->next = temp;
    tail = temp;
    return;
}

string Calculator::getListStr()
{
    stringstream res;
    node *it = head;
    char last = it->C;
    while (it != NULL)
    {
        char c = it->C;
        if (c == '(' && last == ')')
            res << "* ( ";
        else if (isdigit(c) && last == ')')
            res << " * " << c;
        else if (isdigit(last) && c == '(')
            res << " * "
                << "( ";
        else if (c == '(' || c == ')')
            res << ' ' << c << ' ';
        else if (isdigit(c) || (is_operator(last) && is_operator(c)))
            res << c;
        else if (it != cursor)
            res << ' ' << c << ' ';
        last = c;
        it = it->next;
    }
    return res.str();
}

string Calculator::print()
{
    stringstream ss;
    for (node *curr = head; curr != NULL; curr = curr->next)
        ss << curr->C;
    return ss.str();
}

bool Calculator::is_number(const string &s)
{
    bool valid = true;
    string::const_iterator it = s.begin();
    if ((*it == '-' || *it == '+'))
    {
        it++;
        valid = s.length() > 1;
    }
    while (it != s.end() && isdigit(*it))
        ++it;
    return !s.empty() && it == s.end() && valid;
}

bool Calculator::is_operator(char c) { return c == '+' || c == '-' || c == '*'; }

int Calculator::prec(char operand)
{
    switch (operand)
    {
    case '*':
        return 3;
    case '+':
    case '-':
        return 2;
    case '(':
        return 1;
    default:
        return 0;
    }
}

string Calculator::postfix(string exp)
{
    stringstream sinfix(exp);
    stringstream postfix;
    stack<char> opStack;
    string token;
    while (getline(sinfix, token, ' '))
    {
        if (is_number(token))
            if (token.length() > 8)
                postfix << longMod(token) << ' ';
            else
                postfix << token << ' ';
        else if (token[0] == '(')
            opStack.push(token[0]);
        else if (token[0] == ')')
        {
            char top = opStack.top();
            opStack.pop();
            while (top != '(')
            {
                postfix << top << ' ';
                top = opStack.top();
                opStack.pop();
            }
        }
        else if (prec(token[0]))
        {
            while (!opStack.empty() && (prec(opStack.top()) >= prec(token[0])))
            {
                postfix << opStack.top() << ' ';
                opStack.pop();
            }
            opStack.push(token[0]);
        }
    }
    while (!opStack.empty())
    {
        postfix << opStack.top() << ' ';
        opStack.pop();
    }
    return postfix.str();
}

int main()
{
    int n;
    string in;
    cin >> n;
    cin >> in;
    Calculator calculator(in);
    getline(cin, in);
    for (int i = 0; i < n; i++)
    {
        getline(cin, in);
        switch (in[0])
        {
        case '<':
            calculator.moveCursor('<');
            break;
        case '>':
            calculator.moveCursor('>');
            break;
        case '+':
            calculator.insert(in[2]);
            break;
        case '-':
            calculator.remove();
            break;
        case '?':
            cout << calculator.print() << endl;
            break;
        case '!':
            cout << calculator.evaluate() << endl;
            break;
        }
    }
    return 0;
}