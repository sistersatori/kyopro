#include <iostream>
using namespace std;

int main()
{
    int status;
    cin >> status;
    if (status/100 == 4)
        cout << "Yes" << endl;
    else if (status/100 == 5)
        cout << "Yes" << endl;
    else
        cout << "No" << endl;
    return 0;
}
