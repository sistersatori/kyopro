#include <iostream>
using namespace std;

int main()
{
    int n;
    int a[50];
    int ans = 0;

    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> a[i];
        ans += a[i];
    }

    cout << ans << endl;
    return 0;
}
