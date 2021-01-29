import React, {useEffect, useState} from "react";
import {userSession} from "./UserSession";
import Header from "./Components/Header/Header";
import Home from "./Components/Home/Home";
import {HashRouter, Route} from "react-router-dom"
import SupplyList from "./Components/Supply List/SupplyList";
import Account from "./Components/Account/Account";
import Users from "./Components/Users/Users";
import DisplayError from "./Components/Error/DisplayError";
import Navbar from "./Components/NavBar/Navbar";
import Roles from "./Components/Roles/Roles";
import School from "./Components/School/School";
import {canView} from "./Components/Permissions/Permissions";

function App() {
    const [user, setUser] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        // Refresh session every 10 minutes
        setInterval(refreshSession, 600000)
        // Listen to localstorage changes
        // Update user on change
        window.addEventListener('storage', () => {
            setUser(JSON.parse(localStorage.getItem("user")))
        });

        //Fetch user from api
        fetch("/oauth/v1/account")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
            })
            .then(
                (result) => {
                    setUser(result);
                    localStorage.setItem("user", JSON.stringify(result))
                }, (error) => {
                    setUser(null)
                    localStorage.removeItem("user")
                    setError(error);
                }
            )
    }, [])

    function refreshSession() {
        fetch("/oauth/v1/refresh")
            .then((res) => {
                if (!res.ok) {
                    setError("Unable to refresh session")
                }
            })
    }

    return (
        error === null ? (
            <HashRouter>
                <userSession.Provider value={[user, setUser]}>
                    <Header/>
                    <main>
                        <Route exact path="/" component={Home}/>
                        <Route path="/list/:id" component={SupplyList}/>
                        {user != null && (
                            <Route exact path="/account" component={Account}/>
                        )}
                        {canView("user", user) && (
                            <Route exact path="/users" component={Users}/>
                        )}

                        {canView("role", user) && (
                            <Route exact path="/roles" component={Roles}/>
                        )}
                        <Route path="/school/:id" component={School}/>
                    </main>
                    <Navbar/>
                </userSession.Provider>
            </HashRouter>
        ) : (
            <DisplayError msg={"DisplayError!"}/>
        )
    );
}
export default App