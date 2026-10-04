// Chapter 1 content frozen from GIECSER_FINAL (supervisor version):
// goal setting → self-reflection. Self-reflection questions and the
// repractice loop live in the frontend — no rows by design.
package main

import "asri-backend/internal/model"

func ch1Final() seedChapter {
	history := `Riau coastal traditional games have been part of Malay community life for many generations. They were usually played by children, teenagers, and adults during their free time or at community gatherings. Many of these games used simple materials from the local environment, such as wood, coconut shells, bamboo, rattan, seeds, and shells. This shows the close relationship between the Malay people and their natural environment. Traditional games were not only for fun. They also helped people develop important skills and values. Children learned how to cooperate, communicate, follow rules, concentrate, solve problems, and respect other players. One example is Rimau, a traditional strategy game associated with Malay coastal villages. Other traditional games associated with Riau include gasing, congkak, bakiak, layang-layang, and sengki. Riau coastal traditional games are valuable cultural heritage.`
	kinds := `Bakiak Papan congkak Selop Gasing Gasing pangkah Gasing uri Layang layang Jong layo Pacu kolek Lulu Cina Buta Ular naga Selambut Benteng Boi Boian Engrang Statak Galah Panjang`
	jongShort := `Jong Layo is a traditional boat game from Riau, Indonesia. The game uses miniature boats that are usually made and decorated creatively. The boats are placed on the water and can move with the help of wind and water. Jong Layo is not only a fun activity but also reflects the creativity and cultural traditions of the people in Riau.`
	gameWords := `Miniature boat Boat Water Competition Tradition Creativity Game Culture Team Move Win Lose`
	nouns := `Traditional Game Culture Local Player Sail Team Decorate Float Watch Group Partner Opponent Playground Ball Rope Stick Stone Seed Board Line Circle Square`
	verbs := `Play Run Jump Hope Throw Catch Kick Spin Chase Hide Touch Move Count Win Lose Seek`
	adjs := `Fun Exciting Interesting Challenging Easy Difficult Enjoyable Traditional Active Competitive Creative Educational`
	malayNames := `Bakiak Congkak Gasing Jong Layo Galah Panjang Ular Naga Rimau Sengki`
	langFunc := `Asking information: What game is this? How do you play it? How many players do you need? Giving information: It is called Congklak. Two players can play the game. We need a board and small seeds. Explaining how to play: First, prepare the equipment. Next, put the seeds in the holes. Then, choose a player to start. After that, move the pieces. Finally, count the points. Likes: I like Galah Panjang because it is exciting and I can play it with my friends. Inviting: Do you want to play Congklak? Let's play Galah Panjang!`
	grammar := `Simple present: Two players play Congklak. Nominal sentence: Jong Layo is a unique traditional game from Riau. Verbal sentence: The players put the small boat on the water. Imperatives: Stand in a line. Choose a player. Count the points. Modals: You can play this game with your friends. Sequencing: First, Next, Then, After that, Finally.`
	langExpr := `Make: Players make and decorate their own small boats creatively. Sail: The wooden boats sail smoothly when the wind blows. Watch: Children watch the boats float across the pond. Combine: The game combines creativity, tradition, and fun. Traditional: Jong Layo is a popular traditional game in coastal regions of Riau. Unique: The game is unique because it relies entirely on natural wind and water.`
	paiman := `Juminah: Hi! What are you doing? Paiman: Hi! I'm learning about a traditional game from Riau called Jong Layo. Juminah: Jong Layo? What is it? Paiman: Jong Layo is a traditional game that uses a small boat. Juminah: How do you play it? Paiman: First, we prepare and decorate a small boat. Then, we put the boat on the water. Juminah: Does the boat move by itself? Paiman: The boat can move with the wind or water. Juminah: That sounds interesting. Is it difficult to play? Paiman: No, I think it is quite simple and fun. Juminah: What can we learn from this game? Paiman: We can learn creativity and appreciate the traditional culture of Riau.`
	jongLong := `Hello everyone. Today, I want to talk about a traditional game from Riau called Jong Layo. Jong Layo is a traditional game that uses a small boat. The boat is usually made and decorated creatively. Players put the boat on the water and make it move or sail. The game can be played by several people. The players prepare their boats and then put them on the water. They can watch the boats float and move with the wind or water. Jong Layo is an interesting traditional game because it is not only fun but also creative. Players can learn how to make and decorate a small boat. It also helps us appreciate the traditional culture of Riau. I think Jong Layo is a unique game because it combines creativity, tradition, and fun.`
	sarahDaveGap := `Sarah: Hi, Dave! What are you doing? Dave: Hi, Sarah! I'm learning about a traditional game from (1) ________ called Jong Layo. Sarah: Oh, what is Jong Layo? Dave: It is a traditional game that uses a (2) ________. Sarah: How do people play it? Dave: First, the (3) ________ prepare their boats. They also (4) ________ the boats. Sarah: Where do they put the boats? Dave: They put the boats on the (5) ________. Sarah: How do the boats move? Dave: They can move with the (6) ________ or water. Sarah: That sounds interesting! Is the game difficult? Dave: No. It is fun and (7) ________. Sarah: What can people learn from Jong Layo? Dave: They can learn about creativity and appreciate the traditional (8) ________ of Riau.`
	galah := `Galah Panjang is a traditional team game. Players are divided into two teams. One team tries to cross the lines while the other team guards the lines. Players need to run and move quickly. The players cannot be touched by the guards. The game needs teamwork and strategy. It is fun and exciting to play with friends. Galah Panjang can be played by adults or children.`
	tengku := `Tengku: Hey, we have some free time tomorrow afternoon. Do you want to hang out with the kids from the neighborhood? Siti: Sounds good! What do you have in mind? Tengku: How about we get them to play Galah Panjang at the field near the village hall? Siti: Oh, that's a great idea! But do we have enough players? Tengku: Don't worry, they're totally down. We can split into two teams. Siti: Awesome! But get ready to run a lot. Tengku: Exactly, I'll bring the chalk to draw the field lines.`
	pacu := `Pacu Kolek is made from a coconut tree. Pacu Kolek is played in a team that consists of 4 to 10 players.`

	mcq := func(n int, stem, a, b, c, d, correct string) seedTask {
		return quizTask(model.TaskMultipleChoice,
			"Choose the correct answer (A, B, C, or D), then say the complete sentence aloud: "+stem,
			`{"A":"`+a+`","B":"`+b+`","C":"`+c+`","D":"`+d+`"}`,
			`{"correct":"`+correct+`"}`)
	}

	return seedChapter{
		title: "Traditional Games",
		desc:  "Mendeskripsikan dan menjelaskan aturan permainan tradisional Melayu Riau pesisir secara lisan dalam bahasa Inggris.",
		courses: []seedCourse{
			{title: "Goal Setting", desc: "Tujuan pembelajaran yang hendak dicapai.", modules: []seedModule{
				{title: "My Learning Goals", typ: model.ModuleMonologue,
					content:    "By the end of the lesson: I can explain Jong Layo Riau Traditional Games in English for 5-10 minutes correctly and fluently. I can describe how to play Jong Layo Traditional Games in English for 5-10 minutes correctly and fluently.",
					transcript: "I can explain Jong Layo Riau Traditional Games in English for 5 to 10 minutes correctly and fluently. I can describe how to play Jong Layo Traditional Games in English correctly and fluently.",
					order:      1},
			}},
			{title: "Input & Exploration", desc: "Passages, vocabulary and pronunciation models.", modules: []seedModule{
				{title: "History of Coastal Games", typ: model.ModuleMonologue, content: "Read the passage about the history of Riau coastal traditional games.", transcript: history, order: 1, tasks: []seedTask{
					speakTask("Read the history passage aloud with clear pronunciation.", "rimau => ree-mow; gasing => gah-sing; congkak => chong-klahk; bakiak => bah-kee-ahk; sengki => seng-kee"),
				}},
				{title: "Kinds of Games", typ: model.ModuleMonologue, content: "Seventeen traditional games from Table 1: Bakiak, Papan congkak, Selop, Gasing, Layang-layang, Jong layo, Pacu kolek, Lulu Cina Buta, Ular naga, Selambut, Benteng, Boi Boian, Engrang, Statak, Galah Panjang.", transcript: kinds, order: 2, tasks: []seedTask{
					speakTask("Name five games from Table 1 and say what each uses.", "bakiak => bah-kee-ahk; congkak => chong-klahk; selop => seh-lop; layang layang => lah-yahng; pacu kolek => pah-choo koh-lek; ular naga => oo-lahr nah-gah; benteng => behn-teng; engrang => eng-rahng; statak => stah-tahk"),
				}},
				{title: "Jong Layo", typ: model.ModuleMonologue, content: "Listen to the monologue, then read it aloud.", transcript: jongShort, order: 3, tasks: []seedTask{
					speakTask("Read the Jong Layo monologue aloud.", "jong layo => jong lah-yoh"),
				}},
				{title: "Game Words", typ: model.ModuleMonologue, content: "Twelve words with meanings: miniature boat, boat, water, competition, tradition, creativity, game, culture, team, move, win, lose.", transcript: gameWords, order: 4, tasks: []seedTask{
					speakTask("Read each word aloud, then use three in your own sentences.", "miniature => mih-nee-ah-chur; competition => kom-peh-tih-shun; creativity => kree-ah-tih-vih-tee"),
				}},
				{title: "Nouns Pronunciation", typ: model.ModuleMonologue, content: "Listen and repeat each noun.", transcript: nouns, order: 5, tasks: []seedTask{
					speakTask("Read the noun list aloud.", ""),
				}},
				{title: "Verbs Pronunciation", typ: model.ModuleMonologue, content: "Listen and repeat each verb.", transcript: verbs, order: 6, tasks: []seedTask{
					speakTask("Read the verb list aloud.", ""),
				}},
				{title: "Adjectives Pronunciation", typ: model.ModuleMonologue, content: "Listen and repeat each adjective.", transcript: adjs, order: 7, tasks: []seedTask{
					speakTask("Read the adjective list aloud.", ""),
				}},
				{title: "Malay Game Names", typ: model.ModuleMonologue, content: "Say each traditional game name clearly.", transcript: malayNames, order: 8, tasks: []seedTask{
					speakTask("Pronounce each game name clearly.", "bakiak => bah-kee-ahk; congkak => chong-klahk; gasing => gah-sing; jong layo => jong lah-yoh; galah panjang => gah-lah pahn-jang; ular naga => oo-lahr nah-gah; rimau => ree-mow; sengki => seng-kee"),
				}},
				{title: "Language Function", typ: model.ModuleMonologue, content: langFunc, order: 9},
				{title: "Grammar Focus", typ: model.ModuleMonologue, content: grammar, order: 10},
				{title: "Language Expression", typ: model.ModuleMonologue, content: langExpr, order: 11},
			}},
			{title: "Comprehension Check", desc: "Video, dialogue, quizzes and oral questions.", modules: []seedModule{
				{title: "How to Make Jong Layo (Video)", typ: model.ModuleVideo, content: "Watch how to make Jong Layo and how to play it.", media: "https://drive.google.com/uc?export=download&id=1QvjMecLLTMBoWvuDhK0S6paChAQUZwQo", order: 1, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Jong Layo is commonly played by children in _______.", "", `{"answers":["Riau"]}`),
					quizTask(model.TaskOpenEnded, "What materials are used to make Jong Layo?", "", `{"answers":["Tree trunk, plastic, and rope"]}`),
					quizTask(model.TaskOpenEnded, "How is the procedure to play Jong Layo?", "", `{"answers":["Players release their mini boats into the water from a starting line, letting the wind drive them freely, and the boat that reaches the finish line first wins"]}`),
				}},
				{title: "Paiman and Juminah", typ: model.ModuleDialogue, content: "Listen to the dialogue and answer the questions.", transcript: paiman, order: 2, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", "jong layo => jong lah-yoh"),
				}},
				{title: "Jong Layo Vocabulary Quiz", typ: model.ModuleQuiz, content: "Choose the correct answer, then say the complete sentence aloud.", order: 3, tasks: []seedTask{
					mcq(1, "Jong Layo is a traditional ________ from Riau.", "Game", "Food", "Song", "Car", "A"),
					mcq(2, "Jong Layo is a traditional game that uses a small ________.", "Kite", "Boat", "Ball", "Drum", "B"),
					mcq(3, "First, we ________ and decorate a small boat.", "Prepare", "Catch", "Throw", "Push", "A"),
					mcq(4, "Then, we put the boat on the ________.", "Road", "Field", "Water", "Wall", "C"),
					mcq(5, "The boat can move with the ________ or water.", "Wind", "Player", "Song", "Rope", "A"),
					mcq(6, "No, I think it is quite simple and ________.", "Difficult", "Dangerous", "Fun", "Easy", "C"),
					mcq(7, "We can learn ________ from this game.", "Creativity", "Swimming", "Cooking", "Cheating", "A"),
					mcq(8, "Jong Layo is part of the traditional ________ of Riau.", "Culture", "Sport", "Food", "Drink", "A"),
					mcq(9, "The players move ________ when they play the game.", "Together", "Boat", "Traditional", "Local", "A"),
					mcq(10, "Jong Layo ________ a traditional game from Riau.", "Prepare", "Is", "Decorate", "Are", "B"),
				}},
				{title: "Sarah and Dave", typ: model.ModuleDialogue, content: "Fill in the blanks using the word bank, then practice speaking the dialogue aloud.", transcript: sarahDaveGap, order: 4, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Fill blanks (1-8) from: Riau, Water, Wind, Culture, Small boat, Decorate, Creative, Players. Then read the complete dialogue aloud.",
						`{"bank":["Riau","Water","Wind","Culture","Small boat","Decorate","Creative","Players"]}`,
						`{"blanks":["Riau","Small boat","Players","Decorate","Water","Wind","Creative","Culture"]}`),
				}},
				{title: "Listening: Jong Layo", typ: model.ModuleMonologue, content: "Match each beginning (1-5) with its ending (A-E), then pronounce your answers aloud. Source monologue: the Jong Layo presentation.",
					transcript: jongLong, order: 5, tasks: []seedTask{
						// ponytail: dialog audio Drive ID is garbled in PDF extraction — supervisor to confirm before seeding audio_model_url.
						quizTask(model.TaskMatching, "Match beginnings 1-5 with endings A-E.",
							`{"statements":{"1":"Today, the speaker wants to talk about...","2":"Jong Layo is a traditional game that uses...","3":"Players put the small boat on the water to...","4":"Players can learn creativity by learning...","5":"The speaker thinks Jong Layo is unique because it can..."},"endings":{"A":"...how to make and decorate a small boat.","B":"...a traditional game from Riau called Jong Layo.","C":"...combine creativity, tradition, and fun.","D":"...make it move or sail.","E":"...a small boat that is decorated creatively."}}`,
							`{"1":"B","2":"E","3":"D","4":"A","5":"C"}`),
					}},
				{title: "Jong Layo Oral Questions", typ: model.ModuleMonologue, content: "Answer each question orally with clear pronunciation.", transcript: jongLong, order: 6, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "What is the name of the traditional game from Riau?", "", `{"answers":["Jong Layo"]}`),
					quizTask(model.TaskOpenEnded, "What does Jong Layo use?", "", `{"answers":["A small boat"]}`),
					quizTask(model.TaskOpenEnded, "How is the boat usually made and decorated?", "", `{"answers":["Creatively"]}`),
					quizTask(model.TaskOpenEnded, "Where do the players put the boat?", "", `{"answers":["Water"]}`),
					quizTask(model.TaskOpenEnded, "How can the boat move?", "", `{"answers":["Wind or water"]}`),
					quizTask(model.TaskOpenEnded, "How many people can play the game?", "", `{"answers":["2 or more"]}`),
					quizTask(model.TaskOpenEnded, "What can players learn from Jong Layo?", "", `{"answers":["How to make and decorate a small boat"]}`),
					quizTask(model.TaskOpenEnded, "Why is Jong Layo an interesting traditional game?", "", `{"answers":["Because it is a miniature sailing boat race and one of the traditional games of Riau"]}`),
					quizTask(model.TaskOpenEnded, "What does Jong Layo combine?", "", `{"answers":["creativity, tradition, and fun"]}`),
					quizTask(model.TaskOpenEnded, "What does Jong Layo help people appreciate?", "", `{"answers":["The traditional culture of Riau"]}`),
				}},
				{title: "Story Order", typ: model.ModuleMonologue, content: "Rearrange the sentences into the correct order, then speak the complete story aloud.", transcript: "A. The players prepare their boats and put them on the water. B. Jong Layo is a traditional game from Riau. C. Players can learn how to make and decorate a small boat. D. The game can be played by several people. E. Jong Layo combines creativity, tradition, and fun. F. The boat is usually made and decorated creatively.", order: 7, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Rearrange A-F into the correct story order.", `{"sentences":["A","B","C","D","E","F"]}`, `{"order":["B","F","D","A","C","E"]}`),
				}},
			}},
			{title: "Guided Speaking", desc: "Galah Panjang monologue, dialogue and guiding questions.", modules: []seedModule{
				{title: "Galah Panjang", typ: model.ModuleMonologue, content: "Listen to the monologue and answer the guiding questions orally: How to make Galah Panjang? What materials are used? Where can you get the materials?", transcript: galah, order: 1, tasks: []seedTask{
					speakTask("Answer the guiding questions orally and record.", "galah panjang => gah-lah pahn-jang"),
					quizTask(model.TaskOpenEnded, "Complete with words from the text: Galah Panjang is a traditional (1)____ game. Players are divided into (2)____ teams. One team tries to (3)____ the lines. The game needs teamwork and (4)____. It can be played by (5)____ or children.",
						"", `{"blanks":["team","two","cross","strategy","adults"]}`),
				}},
				{title: "Tengku and Siti", typ: model.ModuleDialogue, content: "Practice the dialogue with your partner, then record your voice.", transcript: tengku, order: 2, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", "galah panjang => gah-lah pahn-jang"),
				}},
				{title: "Galah Panjang Vocabulary", typ: model.ModuleMonologue, content: "Miniature boat, competition, tradition, creativity, kolek, race, team, culture, cooperation, prize, winner, river — with example sentences. Sentence starters: Jong Layo is a... It is made from... People usually made Jong Layo for...", order: 3},
			}},
			{title: "Independent Speaking", desc: "Present Pacu Kolek without reading the script.", modules: []seedModule{
				{title: "Pacu Kolek Presentation", typ: model.ModuleOralTest, content: "Speak for 3 to 5 minutes about Pacu Kolek: physical appearance, preparation, benefits.", transcript: pacu, order: 1, tasks: []seedTask{
					speakTask("Present Pacu Kolek for 3 to 5 minutes without reading, then record.", "pacu kolek => pah-choo koh-lek"),
				}},
				{title: "Language Functions Reference", typ: model.ModuleMonologue, content: "Describing: Pacu Kolek is a traditional boat race. Asking: Where is it held? Giving information: It is held in Riau. Opinions: I think it is exciting. Reasons: I like it because it is exciting. Wishes, agreeing, responding, appreciation frames.", order: 2},
			}},
		},
	}
}
