import './Home.scss'
import {useContext, useEffect, useState} from "react";
import SchoolCard from "./SchoolCard";
import {userSession} from "../../UserSession";

function Home() {
    const [schools, setSchools] = useState(null);
    const [user] = useContext(userSession);
    const [error, setError] = useState(null)

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/schools")
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
            })
            .then(
                (result) => {
                    setSchools(result);
                }, (error) => {
                    setSchools(null)
                    setError(error);
                }
            )
    }, [])

    return (
        error === null && schools !== null &&
        <div className="home">
            <section className="schools">
                {schools.map((school) =>
                    <SchoolCard school={school}/>
                )}
            </section>

            {user != null &&
            user.roles.consolidated_roles.resources.schools.policy.can_add && (
                <div>
                    Create School
                    <input placeholder="name"/>
                    <button>Create</button>
                </div>
            )
            }
        </div>
    );

}

export default Home;
