// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(Debug)]
pub struct Duration {
    seconds: u64,
}

const EARTH_YEAR_IN_SECONDS: u64 = 31_557_600;

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration { seconds: s }
    }
}

pub trait ToSeconds {
    fn to_seconds(&self) -> u64;
}

impl ToSeconds for Duration {
    fn to_seconds(&self) -> u64 {
        self.seconds
    }
}

pub trait Planet {
    fn relative_orbital_period() -> f64;
    fn years_during(d: &Duration) -> f64 {
        let base_years = d.to_seconds() as f64 / EARTH_YEAR_IN_SECONDS as f64;
        base_years / Self::relative_orbital_period()
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

macro_rules! planet {
    ($planet:ident, $orbital_period:expr) => {
        impl Planet for $planet {
            fn relative_orbital_period() -> f64 {
                $orbital_period
            }
        }
    };
}

planet!(Earth, 1.0);
planet!(Mercury, 0.2408467);
planet!(Venus, 0.61519726);
planet!(Mars, 1.8808158);
planet!(Jupiter, 11.862615);
planet!(Saturn, 29.447498);
planet!(Uranus, 84.016846);
planet!(Neptune, 164.79132);
