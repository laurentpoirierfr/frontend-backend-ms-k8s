import logo from './logo.svg';
import './App.css';
import Navbar from './Components/Navbar';
import About from './Pages/About';
import Login from './Pages/Login';
import Home from './Pages/Home';
import Skills from './Pages/Skiils';
import Projects from './Pages/Projects';


function App() {
  return (
    <div className="App">
      <Navbar />
      <Home/>
      <About />
      <Skills />
      <Projects/>
      <Login/>
    </div>
  );
}

export default App;
