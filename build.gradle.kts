plugins {
    kotlin("jvm") version "1.9.0"
    application
    id("com.github.johnrengelman.shadow") version "8.1.1"
}

group = "dev.apachejuice"
version = "0.1"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation(kotlin("test"))
}

tasks.test {
    useJUnitPlatform()
}

tasks.shadowJar {
    archiveFileName.set("${project.group}-${project.name}-$version.jar")
}

kotlin {
    jvmToolchain(8)
}

application {
    mainClass.set("dev.apachejuice.pretzel.compiler.MainKt")
}