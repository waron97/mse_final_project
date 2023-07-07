export default function splitPassages(text: string): string[] {
  const maxLength = 64;
  const regex = /\(?[^\.\?\!]+[\.!\?]\)?/g;
  let sentences = text.match(regex);
  let passages: string[] = [];
  if (!sentences) {
    return passages;
  } else {
    let currLength = 0;
    let currPassage = "";
    let i = 0;
    while (i < sentences.length) {
      let tokens = sentences[i].split(" ");
      if (currLength + tokens.length <= maxLength) { // combines several short sentences together in one passage
        currPassage = currPassage.concat(sentences[i]);
        currLength += tokens.length;
      } else { // if adding the next sentences exceed the maxLength
        if (currPassage != "") {  // save the current passage (combination of the previous short sentences) in passages
          passages.push(currPassage);
        }

        if (tokens.length <= maxLength) { // if the next sentence is shorter than maxLength, push the next sentence into passages
          passages.push(sentences[i]);
        } else { // if the next sentence is longer than maxLength, split the next sentence
          for (let j = 0; j < tokens.length; j += maxLength) {
            currPassage = tokens.slice(j, j + maxLength).join(" ");
            if (currPassage != "") {
              passages.push(currPassage);
            }
          }
        }
        currPassage = "";
        currLength = 0;
      }
      i++;
    }

    // add rest of the sentences to the passage
    if (currPassage != "") {
      passages.push(currPassage);
    }
    return passages;
  }
} 
