import Inferno from 'inferno';

export function App({ children }) {
    return (
        <div>
            <h1>Password Generator</h1>
            <p>Send us a username and we'll create a secure password for you!</p>
            { children }
        </div>
    );
}
