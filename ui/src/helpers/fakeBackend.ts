export function configureFakeBackend() {
    let users = [{ id: 1, email: 'test@julian.pro', password: 'testpassword' }]
    let realFetch = window.fetch
    // @ts-ignore
    window.fetch = function (url: string, opts?: any): Promise<Response> {
        let isLoggedIn: boolean = false
        if (opts && opts.headers) {
            isLoggedIn = opts.headers['Authorization'] === 'Bearer fake-jwt-token'
        }

        return new Promise((resolve, reject) => {
            // wrap timeout to simulate server api call
            setTimeout(() => {
                if (url.endsWith('/users/authenticate') && opts.method === 'POST') {
                    const params = JSON.parse(opts.body);
                    const user = users.find(x => x.email === params.email && x.password === params.password);
                    if (!user) return error('Username or password is incorrect')
                    return ok({
                        id: user.id,
                        email: user.email,
                        token: 'fake-jwt-token'
                    });
                }

                // get users - secure
                if (url.endsWith('/users') && opts.method === 'GET') {
                    if (!isLoggedIn) return unauthorised();
                    return ok(users);
                }

                // pass through any requests not handled above
                realFetch(url, opts).then(response => resolve(response));

                // private helper functions
                function ok(body: any) {
                    resolve({
                        ok: true,
                        text: () => Promise.resolve(JSON.stringify(body))
                    } as Response)
                }

                function unauthorised() {
                    resolve({ status: 401, text: () => Promise.resolve(JSON.stringify({ message: 'Unauthorised' })) } as Response)
                }
                function error(message: string) {
                    resolve({ status: 400, text: () => Promise.resolve(JSON.stringify({ message })) } as Response)
                }
            }, 500)
        })
    } 
}