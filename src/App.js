import "./App.css";
import Login from "./component/Login";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Register from "./component/Register";
import Homepage from "./component/Homepage";

function App() {
  return (
   
    <Router>
      
      <Routes>
        {/*   Di dalam Routes yang kita akan tambahlan endpoit - endpoint nya  */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/homepage"  element={<Homepage />} />   
      </Routes>
    </Router>
  );
}

export default App;
