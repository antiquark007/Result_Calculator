import Calculator from "../screens/Calculator";
import CgpaCalculator from "../screens/CgpaCalculator";
import Home from '../screens/Home';

const Router = [
  //home page
  {
    path: "/", element: <Home />
  },
  //college
  {
    path: "/dsce",
    element: <Home />,
  },
  // calculator
  {
    path: "/:college/:branch/:semester/calculator",
    element: <Calculator />,
  },

  // cgpa calculator
  {
    path: "/:college/:branch/:graduatingYear/cgpa",
    element: <CgpaCalculator />,
  },
]

export default Router;