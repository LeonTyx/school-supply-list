import './Header.scss'
import LoginButton from "./LoginButton";
import {Link} from "react-router-dom";

export default function Header() {
    return (
        <header>
            <Link to="/" className="site-name">School Supply Lists</Link>
            <LoginButton/>
        </header>
    );

}
