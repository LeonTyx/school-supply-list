import React, {useEffect, useState} from "react";
import {userSession} from "./UserSession";
import Header from "./Header/Header";
import Home from "./Home";
import {HashRouter, Route} from "react-router-dom"


function App() {
    const [user, setUser] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        setInterval(refreshSession, 900000)
        window.addEventListener('storage', () => {
            setUser(JSON.parse(localStorage.getItem("user")))
        });

        fetch("/oauth/v1/profile")
            .then((res) => {
                if(res.ok){
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
            .then(res => res.json())
            .then(
                (error) => {
                    setError(error);
                }
            )
    }

    return (
        <HashRouter>
            <userSession.Provider value={[user, setUser]}>
                <Header/>
                <main>
                    <Route exact path="/" component={Home}/>
                    <Route exact path="/list/:id" component={Home}/>
                </main>
            </userSession.Provider>
        </HashRouter>
    );
}

export default App