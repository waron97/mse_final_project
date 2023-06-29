import { Frontier } from "../../db/frontier";

const initialPages = [
  {
    url: "https://www.tuebingen.de/",
    priority: 1,
  },
  {
    url: "https://www.tuebingen-info.de/",
    priority: 2,
  },
  {
    url: "https://de.wikipedia.org/wiki/T%C3%BCbingen",
    priority: 3,
  },
  {
    url: "https://uni-tuebingen.de/",
    priority: 4,
  },
  {
    url: "https://www.kreis-tuebingen.de/Startseite.html",
    priority: 5,
  },
  {
    url: "https://www.neckarcamping.de/",
    priority: 6,
  },
  {
    url: "https://www.tuemarkt.de/",
    priority: 7,
  },
  {
    url: "https://www.swtue.de/",
    priority: 8,
  },
];

export default async function initFrontier() {
  const pages = await Frontier.getList();
  if (pages.length !== 0) {
    // frontier already initialized
    return;
  }

  for (const page of initialPages) {
    await Frontier.create(page);
  }
}
