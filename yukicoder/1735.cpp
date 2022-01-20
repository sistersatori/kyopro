#include <iostream>
using namespace std;

int main()
{
    string S;
    cin >> S;

    if (S == "E")
        cout << "F" << endl;
    else if (S == "B")
        cout << "C" << endl;
    else
        cout << S << "#" << endl;
    return 0;
}
