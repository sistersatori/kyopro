#include <iostream>
using namespace std;

int main()
{
    int a, b;
    int sock = 1;
    int ans = 0;

    cin >> a >> b;

    while (sock < b) {
        sock = (sock-1) + a;
        ans++;
    }

    cout << ans << endl;

    return 0;
}

