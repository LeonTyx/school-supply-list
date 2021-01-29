import './Home.scss'
import {useEffect, useState} from "react";
import SchoolCard from "./SchoolCard";
import CreateSchool from "./CreateSchool";

function Home() {
    const [schools, setSchools] = useState(null);
    const [error, setError] = useState(null)
    function handleErrors(response, errorMessage) {
        if (!response.ok) {
            setError(errorMessage)
        }
        return response.json();
    }

    useEffect(() => {
        //Fetch user from api
        fetch("/api/v1/schools")
            .then((resp) => handleErrors(resp, "Unable to retrieve roles"))
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

            <CreateSchool/>
        </div>
    );

}

export default Home;
