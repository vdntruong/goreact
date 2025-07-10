import React, { useState } from 'react';
import ReactDOM from 'react-dom/client';
import './Greeting.js'; // Just to show importing another module

const App = () => {
    const [count, setCount] = useState(0);

    return (
        <div className="container">
            <h1>Hello from Simple React UI!</h1>
            <p>This is a basic React app served as static files.</p>
            <div className="counter">
                <p>You clicked the button {count} times.</p>
                <button onClick={() => setCount(count + 1)}>Click Me!</button>
            </div>
            <p className="footer">
                Built with React and Parcel. Current time: {new Date().toLocaleTimeString()}
            </p>
        </div>
    );
};

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);
