import './Home.scss'
import {useEffect, useState} from "react";

function Home() {
    const [schools, setSchools] = useState(null);
    const [error, setError] = useState(null)

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/schools")
            .then((res) => {
                if(res.ok){
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
        <div>
            {schools.map((school) =>
                <div>{school.name}</div>
            )}
        </div>
    );

}

export default Home;
