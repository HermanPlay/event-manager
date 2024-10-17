<script lang="ts">
	let teamMembers = [
		{
			name: 'Filip Shramko',
			role: 'Lead Developer | CEO | CTO | COO | The best guy ever',
			description:
				'Filip is an experienced full-stack developer, has founded 5 start-ups, 6 of them failed, with a passion for building efficient and scalable applications.',
			image: '/filip-shramko.jpg'
		}
	];
</script>

<div class="min-h-screen bg-gray-100 p-8">
	<div class="container mx-auto bg-white shadow-lg rounded-md p-6">
		<h1 class="text-4xl font-bold mb-6 text-center">About Us</h1>
		<p class="text-lg text-gray-700 text-center mb-10">
			Welcome to our web application! We are a dedicated team of developers, designers, and project
			managers working together to create seamless digital experiences. Our mission is to deliver
			high-quality solutions that meet the needs of our users.
		</p>

		<h2 class="text-3xl font-semibold mb-4 text-center">Meet the Team</h2>
		<div class="grid grid-cols-1 gap-6">
			{#each teamMembers as member}
				<div class="p-4 bg-gray-50 rounded-lg shadow-sm text-center">
					<img
						src={member.image}
						alt={member.name}
						class="mx-auto rounded-full mb-4 shadow-md"
						style="height: 450px;"
					/>
					<h3 class="text-xl font-semibold">{member.name}</h3>
					<p class="text-gray-500">{member.role}</p>
					<p class="text-gray-700 mt-2">{member.description}</p>
				</div>
			{/each}
		</div>
	</div>
</div>
