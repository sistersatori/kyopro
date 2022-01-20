#include <iostream>
using namespace std;

int main()
{
    string N;
    cin >> N;

    bool f4{false}, f6{false};

    for (int i = 0; i < N.length(); i++) {
        if(N.substr(i, 1) == "4")
            f4 = true;
        if(N.substr(i, 1) == "6")
            f6 = true;
    }

    if (f4 && f6)
        cout << "Beautiful" << endl;
    else
        cout << "..." << endl;

    return 0;
}
