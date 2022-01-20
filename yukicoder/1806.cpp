#include <iostream>
using namespace std;

string turn(string direction) {
    if (direction == "N")
        return "E";
    else if (direction == "E")
        return "S";
    else if (direction == "S")
        return "W";
    else if (direction == "W")
        return "N";
}

int main()
{
    string A, B;
    cin >> A >> B;

    int ans = 0;
    while (A != B) {
        A = turn(A);
        ans++;
    }

    cout << ans << endl;
    return 0;
}
