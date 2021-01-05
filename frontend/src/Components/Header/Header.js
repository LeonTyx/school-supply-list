import './Header.scss'
import LoginButton from "./LoginButton";

export default function Header() {
    return (
        <header>
            <div className="site-name">School Supply Lists</div>
            <LoginButton/>
        </header>
    );

}
