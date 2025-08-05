"use client"
import Beams from "./beams";
import { Source_Code_Pro } from "next/font/google"
import "./globals.css";

const codeFont = Source_Code_Pro({subsets: ["latin"]})

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <style>
          @import url('https://fonts.googleapis.com/css2?family=Bitcount+Prop+Single:wght@100..900&family=Google+Sans+Code:ital,wght@0,300..800;1,300..800&display=swap');
        </style>
        <div className="w-screen h-screen absolute z-0 bg-black">
          <Beams
            beamWidth={1}
            beamHeight={25}
            beamNumber={50}
            lightColor="rgb(191, 81, 251)"
            speed={1}
            noiseIntensity={1.5}
            scale={0.2}
            rotation={60}
          />
        </div>
        <div id="mainContainer" className="w-screen h-screen absolute z-10 max-w-full max-h-full overflow-hidden codeFont text-xl">
          {children}
        </div>
      </body>
    </html>
  );
}
