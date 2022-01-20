#include <iostream>
using namespace std;

long long fact(int n)
{
    if (n < 2)
        return n;
    else
        return n * fact(n - 1);
}

int main()
{
    int trick, treat;
    cin >> trick >> treat;
    cout << fact((trick | treat)) << endl;
    return 0;
}
