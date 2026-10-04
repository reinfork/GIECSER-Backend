// Chapter 3 content frozen from GIECSER_FINAL Lesson 3 (pp.45-63):
// goal setting → independent speaking. Self-reflection questions and the
// repractice loop live in the frontend — no rows by design.
package main

import "asri-backend/internal/model"

func ch3Final() seedChapter {
	riverPassage := `Riau is a region famously shaped by its vast network of winding rivers, coastal areas, and long shorelines. Throughout history, water transportation has been the absolute backbone of daily life for local communities. Long before modern infrastructure existed, people relied entirely on traditional transportation to connect with neighboring villages and towns. To navigate both calm inland streams and the open sea, people built various types of boat and traditional boat designs tailored to their needs. For traveling across narrow river branches, a simple wooden canoe or a small sampan was commonly used. Meanwhile, larger passenger boat vessels and even a makeshift ferry helped groups of people cross wider rivers from one dock to another. Every journey required essential gear. A boatman or a skilled fisherman would expertly handle a wooden oar or a dual-bladed paddle to maneuver through the water. For longer trips, some vessels hoisted a large sail to catch the wind, while modern motorized boats rely on a mechanical engine to travel faster. When docking or securing the watercraft, thick rope was always kept on board. Today, whether it is a small fishing boat braving the tides or a busy local vessel carrying every passenger, these watercraft remain a proud symbol of Riau's rich maritime heritage.`
	vessels := `Perahu cadik Jongkong Kayak Sampan/bidak Pompong Sampan layar Sampan leper Jelatik Perahu pencalang Sampan kotok Boat Ferry Wooden fishing rowing sailing motorized passenger`
	vocabList := `Transportation Traditional boat Canoe Ferry Fishing Passenger River Sea Coastal Shore Dock Boatman Oar Paddle Engine Rope Sail`
	jongkongMono := `Hello everyone. Today, I want to talk about a traditional wooden canoe from Riau called the jongkong. A jongkong is a small, lightweight boat carved out from a single, massive tree trunk. Long before modern bridges and roads existed, local people relied on this traditional boat to travel along winding rivers, cross narrow streams, and go fishing. To use it, a boatman simply steps inside and uses a paddle to glide smoothly across the water. Today, even though modern transportation is everywhere, the jongkong remains a wonderful symbol of our rich local heritage.`
	jamesKevin := `Kevin: "Hi, James! What are you reading?" James: "I am reading about a traditional boat from Riau called jongkong." Kevin: "Oh, the small wooden boat? How do people make it?" James: "They carve it out from a single, massive tree trunk." Kevin: "That is amazing! How does a boatman move it on the water?" James: "He uses a wooden paddle to paddle and glide smoothly." Kevin: "Are there still people using it today?" James: "Yes, even though modern transport is everywhere, it remains an important symbol of our local culture."`
	raniMono := `Hello, everyone. My name is Rani. Today, I would like to talk about transportation in the Riau Islands. The Riau Islands have many islands, so people need boats to travel from one island to another. One type of boat that people use is called a pompong. A pompong is a small motorized boat commonly used in coastal areas. People use pompong to go to nearby islands, visit relatives, go to school, go to work, or carry goods. They usually take a pompong from a small jetty or harbor. Travelling by pompong can be an interesting experience. We can see the beautiful sea, small islands, and other boats during the journey. However, passengers need to be careful. They should get on and off the boat carefully and follow the boat driver's instructions. For people who live in coastal areas and small islands, boats are important transportation. Pompong is not only a means of transportation but also part of the daily life and maritime culture of the people in the Riau Islands. Thank you for listening.`
	pompongDialog := `Alya: Good morning, Rina. Where are you going? Rina: Good morning, Alya. I am going to Penyengat Island. Alya: How are you going there? Rina: I am going by pompong. Alya: Really? I have never ridden a pompong before. Rina: It is fun. Pompong is a traditional boat used by people in the Riau Islands. Alya: Where can we take the pompong? Rina: We can take it from the jetty. Come with me! Alya: Okay. How long does the trip take? Rina: It takes about fifteen minutes. Alya: Look! Is that our pompong? Rina: Yes, it is. Let's get on. Mr. Ahmad: Good morning. Where are you going? Rina: We are going to Penyengat Island, Sir. Mr. Ahmad: All right. Please get on carefully. Alya: Thank you, Sir. Mr. Ahmad: Please sit down and hold the boat firmly. Alya: The sea is beautiful! Rina: Yes. We can see many small islands from the boat. Alya: I like travelling by pompong. It is an important part of life in the Riau Islands. Rina: That's right. Many people use boats to travel between islands. Alya: I hope I can ride a pompong again. Rina: Me too!`

	mcq := func(stem, a, b, c, d, correct string) seedTask {
		return quizTask(model.TaskMultipleChoice,
			"Choose the correct answer (A, B, C, or D), then say the complete sentence aloud: "+stem,
			`{"A":"`+a+`","B":"`+b+`","C":"`+c+`","D":"`+d+`"}`,
			`{"correct":"`+correct+`"}`)
	}

	return seedChapter{
		title: "Maritime Transport",
		desc:  "Memahami informasi tentang transportasi tradisional perairan Riau dan menceritakannya secara lisan.",
		courses: []seedCourse{
			{title: "Goal Setting", desc: "Tujuan pembelajaran yang hendak dicapai.", modules: []seedModule{
				{title: "My Learning Goals", typ: model.ModuleMonologue,
					content:    "By the end of the lesson: I can explain about how to ride ponpon in English for 5-10 minutes correctly and fluently. I can describe my experience traveling using Maritim Transportatation in English for 5-10 minutes correctly and fluently.",
					transcript: "I can explain about how to ride ponpon in English for 5 to 10 minutes correctly and fluently. I can describe my experience traveling using maritime transportation in English correctly and fluently.",
					order:      1},
			}},
			{title: "Input & Exploration", desc: "Passages, vessels and pronunciation models.", modules: []seedModule{
				{title: "Riau Maritime Heritage", typ: model.ModuleMonologue, content: "Read the passage about Riau Malay maritime transportation.", transcript: riverPassage, order: 1, tasks: []seedTask{
					speakTask("Read the passage aloud with clear pronunciation.", "infrastructure => in-fruh-struk-chur; maneuver => muh-noo-ver; watercraft => waw-ter-kraft"),
				}},
				{title: "Riau Malay Vessels", typ: model.ModuleMonologue, content: "Eighteen vessels from Table 3.1: Perahu cadik, Jongkong, Kayak, Sampan/bidak, Pompong, Sampan layar, Sampan leper, Jelatik, Perahu pencalang, Sampan kotok, Boat, Ferry, Wooden boat, Fishing boat, Rowing boat, Sailing boat, Motorized boat, Passenger boat.", transcript: vessels, order: 2, tasks: []seedTask{
					speakTask("Name five traditional vessels and describe their build.", "jongkong => jong-kong; pompong => pom-pong; cadik => chah-dik; jelatik => jeh-lah-tik; pencalang => pen-chah-lahng"),
				}},
				{title: "Maritime Vocabulary", typ: model.ModuleMonologue, content: "Pronunciation model words: Transportation, Boat, Canoe, Sea, Fisherman, Engine, Rope, Sail, Dock, Boatman, Oar, Paddle.", transcript: vocabList, order: 3, tasks: []seedTask{
					speakTask("Read each vocabulary word aloud clearly.", "transportation => trans-por-tay-shun; fisherman => fish-er-mun; engine => en-jin; paddle => pad-ul"),
				}},
				{title: "Jongkong Canoe", typ: model.ModuleMonologue, content: "Listen to the monologue about Jongkong, then read it aloud.", transcript: jongkongMono, order: 4, tasks: []seedTask{
					speakTask("Read the Jongkong monologue aloud.", "jongkong => jong-kong; lightweight => lite-wayt; trunk => trungk"),
				}},
				{title: "James and Kevin Dialogue", typ: model.ModuleDialogue, content: "Practice the dialogue about Jongkong with a partner.", transcript: jamesKevin, order: 5, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", "jongkong => jong-kong; smooth => smooth"),
				}},
			}},
			{title: "Comprehension Check", desc: "Quizzes, matching and dialogue completion.", modules: []seedModule{
				{title: "Maritime Vocabulary Quiz", typ: model.ModuleQuiz, content: "Choose the correct answer, then say the complete sentence aloud.", order: 1, tasks: []seedTask{
					mcq("A traditional Riau boat carved out from a single, massive tree trunk is a...", "ferry", "jongkong", "passenger boat", "fishing boat", "B"),
					mcq("A wooden or plastic blade with a long handle used to push a boat manually is a...", "Engine", "Rope", "Sail", "Paddle", "D"),
					mcq("A large cloth attached to a mast to catch the wind without an engine is a...", "sail", "dock", "shore", "ferry", "A"),
					mcq("The land along the edge of a sea, lake, or large river is known as the...", "river", "shore", "engine", "passenger", "B"),
					mcq("A boat designed specifically to carry people across water on regular routes is a...", "fishing boat", "passenger boat", "canoe", "paddle", "B"),
					mcq("A mechanical machine that makes motorized boats move fast is the...", "paddle", "rope", "engine", "dock", "C"),
					mcq("A strong, thick cord made of twisted fibers used for securing a boat is a...", "sail", "rope", "paddle", "shore", "B"),
					mcq("An area by the water where boats arrive and passengers get on or off is a...", "dock", "sea", "coastal area", "river", "A"),
					mcq("A person whose job is to catch fish in rivers or the sea is a...", "boatman", "passenger", "fisherman", "ferry", "C"),
					mcq("The broad, salty body of water covering most of the Earth's surface is the...", "river", "sea", "dock", "Shore", "B"),
				}},
				{title: "Listening: Jongkong", typ: model.ModuleMonologue, content: "Match each beginning (1-5) with its ending (A-E), then pronounce your answers aloud. Source monologue: the Jongkong presentation.", transcript: jongkongMono, order: 2, tasks: []seedTask{
					// ponytail: corrected statement 3/4 mapping per user choice to fix book typo.
					quizTask(model.TaskMatching, "Match beginnings 1-5 with endings A-E.",
						`{"statements":{"1":"A jongkong is a traditional wooden canoe from Riau that is...","2":"Local people in the past relied on the jongkong to travel along winding rivers, cross narrow streams, and...","3":"Long before modern bridges and roads existed, people used the jongkong because...","4":"Players can learn creativity by learning...","5":"Today, the jongkong remains..."},"endings":{"A":"...glide smoothly across the water.","B":"...carved out from a single, massive tree trunk.","C":"...go fishing.","D":"...a wonderful symbol of our rich local heritage.","E":"...there was no modern transportation."}}`,
						`{"1":"B","2":"C","3":"E","4":"A","5":"D"}`),
				}},
				{title: "Jongkong Oral Questions", typ: model.ModuleMonologue, content: "Answer each question orally with clear pronunciation.", transcript: jongkongMono, order: 3, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "What is the main topic of the speaker's presentation?", "", `{"answers":["A traditional wooden canoe from Riau called the jongkong"]}`),
					quizTask(model.TaskOpenEnded, "What specific region is the jongkong from?", "", `{"answers":["Riau"]}`),
					quizTask(model.TaskOpenEnded, "How would you describe the size and weight?", "", `{"answers":["Small and lightweight"]}`),
					quizTask(model.TaskOpenEnded, "What unique method is used to build a jongkong?", "", `{"answers":["Carved out from a single, massive tree trunk"]}`),
					quizTask(model.TaskOpenEnded, "What were the main purposes of using a jongkong in the past?", "", `{"answers":["To travel along winding rivers, cross narrow streams, and go fishing"]}`),
					quizTask(model.TaskOpenEnded, "What specific tool does a boatman use to move it?", "", `{"answers":["A paddle"]}`),
					quizTask(model.TaskOpenEnded, "How does the boat move across the water?", "", `{"answers":["It glides smoothly across the water"]}`),
					quizTask(model.TaskOpenEnded, "What is everywhere in modern times according to the text?", "", `{"answers":["Modern transportation"]}`),
					quizTask(model.TaskOpenEnded, "Despite modern transport, what remains of the jongkong today?", "", `{"answers":["A wonderful symbol of rich local heritage"]}`),
					quizTask(model.TaskOpenEnded, "Why is the jongkong still considered important today?", "", `{"answers":["Because it remains a wonderful symbol of rich local heritage"]}`),
				}},
				{title: "Story Order & Dialogue", typ: model.ModuleMonologue, content: "Rearrange sentences A-E and fill in the dialogue blanks.", transcript: jongkongMono, order: 4, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Rearrange sentences into correct order: 1. To use it, a boatman steps inside and uses a paddle. 2. Hello everyone. Today I want to talk about jongkong. 3. Thank you for listening. 4. A jongkong is a small lightweight boat carved out from a single tree trunk. 5. Today, even though modern transport is everywhere, the jongkong remains a wonderful symbol.",
						`{"sentences":["1","2","3","4","5"]}`, `{"order":["2","4","1","5","3"]}`),
					quizTask(model.TaskOpenEnded, "Fill blanks (1-5) from James & Kevin dialogue: Jongkong, Wooden, Trunk, Paddle, Culture.",
						`{"bank":["Jongkong","Wooden","Trunk","Paddle","Culture"]}`,
						`{"blanks":["Jongkong","Wooden","Trunk","Paddle","Culture"]}`),
				}},
			}},
			{title: "Guided Speaking", desc: "Pompong dialogue and Rani monologue.", modules: []seedModule{
				{title: "Alya and Rina (Pompong)", typ: model.ModuleDialogue, content: "Practice the dialogue about going to Penyengat Island by pompong.", transcript: pompongDialog, order: 1, tasks: []seedTask{
					speakTask("Choose a role and read the pompong dialogue aloud.", "pompong => pom-pong; jetty => jeh-tee; passenger => pas-en-jer; firmly => ferm-lee"),
				}},
				{title: "Transportation in Riau Islands", typ: model.ModuleMonologue, content: "Read Rani's monologue about pompong and island transport.", transcript: raniMono, order: 2, tasks: []seedTask{
					speakTask("Read Rani's monologue aloud clearly.", "relatives => rel-uh-tivz; harbor => hahr-ber; passengers => pas-en-jerz"),
				}},
			}},
			{title: "Independent Speaking", desc: "Present a maritime vessel without reading.", modules: []seedModule{
				{title: "Maritime Vessel Presentation", typ: model.ModuleOralTest, content: "Choose one Riau maritime vessel (e.g. Jongkong or Pompong). Speak for 5 to 10 minutes without reading: physical appearance, preparation, and benefits. Language functions: Describing (A pompong is a small motorized boat). Talking about use (People use pompong to travel between islands). Opinions and reasons (I think travelling by pompong is interesting because we can see the beautiful sea).", transcript: raniMono, order: 1, tasks: []seedTask{
					speakTask("Present your chosen maritime vessel for 5 to 10 minutes without reading, then record.", "pompong => pom-pong; jongkong => jong-kong; motorized => mo-to-ri-zd"),
				}},
			}},
		},
	}
}
