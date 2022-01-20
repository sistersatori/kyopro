#include <iostream>
using namespace std;

int main()
{
    int a, b, c, d, m;
    int ans = -1;
    cin >> a >> b >> c >> d >> m;
    for(int i = a; i <= b; i++) {
        for(int j = c; j <= d; j++) {
            ans = max(ans, (i+j)%m);
        }
    }
    cout << ans << endl;
    return 0;
}
