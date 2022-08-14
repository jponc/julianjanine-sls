import { Route, Routes } from "react-router-dom";
import { CeremonyScreen } from "./screens/CeremonyScreen";
import { ContactUsScreen } from "./screens/ContactUs";
import { FAQScreen } from "./screens/FAQScreen";
import { HomeScreen } from "./screens/HomeScreen";
import { OurStoryScreen } from "./screens/OurStoryScreen";
import { ReceptionScreen } from "./screens/ReceptionScreen";
import { SuppliersScreen } from "./screens/SuppliersScreen";

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomeScreen />} />
      <Route path="/our-story" element={<OurStoryScreen />} />
      <Route path="/ceremony" element={<CeremonyScreen />} />
      <Route path="/reception" element={<ReceptionScreen />} />
      <Route path="/suppliers" element={<SuppliersScreen />} />
      <Route path="/faq" element={<FAQScreen />} />
      <Route path="/contact-us" element={<ContactUsScreen />} />
    </Routes>
  );
}

export default App;
