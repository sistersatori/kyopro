#include <iostream>
using namespace std;

int main()
{
    int X, Y;
    int ans = 0;
    cin >> X >> Y;

    if (Y - X <= 0) ans = 0;
    else {
        ans += (Y - X) / 10;
        if ((Y - X) % 10 > 0) ans++;
    }

    cout << ans << endl;
    return 0;
}
