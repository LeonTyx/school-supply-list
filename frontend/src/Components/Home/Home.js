import './Home.scss'
import {useEffect, useState} from "react";
import SchoolCard from "./SchoolCard";

function Home() {
    const [schools, setSchools] = useState(null);
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
            {schools.map((school) =>
                <SchoolCard school={school}/>
            )}
        </div>
    );

}

export default Home;
