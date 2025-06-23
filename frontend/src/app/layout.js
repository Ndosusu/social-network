import Beams from "./beams";
import "./globals.css";

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <div className="w-screen h-screen absolute z-0">
          <Beams
            beamWidth={1}
            beamHeight={25}
            beamNumber={50}
            lightColor="#BF51FB"
            speed={1}
            noiseIntensity={1.5}
            scale={0.2}
            rotation={60}
          />
        </div>
        <div id="mainContainer" className="w-screen h-screen absolute z-10 max-w-full max-h-full overflow-scroll">
          {children}
        </div>
      </body>
    </html>
  );
}
